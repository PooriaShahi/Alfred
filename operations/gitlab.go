package operations

import (
	"alfred/helpers"
	"alfred/templates"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var client http.Client

type GitlabData struct {
	GitlabUrl      string
	GitlabApiToken string
}

func GetGitlabDataObject() GitlabData {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found !!!")
		}
	}

	url := viper.GetString("gitlab.url")
	if url == "" {
		helpers.GitlabUrlNotFoundErrorHandler()
		url = "https://gitlab.com"
	}

	apiToken := viper.GetString("gitlab.api_token")
	if apiToken == "" {
		helpers.GitlabApiTokenNotFoundErrorHandler()
	}

	g := GitlabData{GitlabUrl: url, GitlabApiToken: apiToken}
	return g
}

func (g *GitlabData) gitlabGetProjectId(gitlabProject string) (string, error) {
	var data []map[string]interface{}
	var projectId string
	Url := g.GitlabUrl + "/api/v4/search?scope=projects&search=" + gitlabProject
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("PRIVATE-TOKEN", g.GitlabApiToken)

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if res.StatusCode == 401 {
		helpers.UnauthorizedGitlabApiToken()
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(result, &data)
	if err != nil {
		return "", err
	}

	for _, v := range data {
		if v["name"] == gitlabProject {
			projectId = fmt.Sprintf("%v", v["id"])
		}
	}
	if projectId == "" {
		fmt.Println("Cannot find project")
		os.Exit(1)
	}
	return projectId, nil
}

func (g *GitlabData) gitlabVariablesPostRequest(projectId string, postBody []byte) ([]byte, error) {
	url := g.GitlabUrl + "/api/v4/projects/" + projectId + "/variables"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", g.GitlabApiToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *GitlabData) gitlabVariablesDeleteRequest(projectId string, key string, env string) ([]byte, error) {
	url := g.GitlabUrl + "/api/v4/projects/" + projectId + "/variables/" + key + "?filter[environment_scope]=" + env
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", g.GitlabApiToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *GitlabData) gitlabGetFileRequest(projectId string, filePath string, branch string) ([]byte, error) {
	var resultData map[string]interface{}
	filePath = strings.ReplaceAll(filePath, "/", "%2F")
	filePath = strings.ReplaceAll(filePath, ".", "%2E")
	url := g.GitlabUrl + "/api/v4/projects/" + projectId + "/repository/files/" + filePath + "?ref=" + branch
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", g.GitlabApiToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(result, &resultData)
	decodedData, _ := base64.StdEncoding.DecodeString(resultData["content"].(string))
	return decodedData, nil
}

func (g *GitlabData) GenerateEnvsSpringBoot(gitlabProject, gitlabBranch string) {
	var yamlFile map[string]interface{}

	projectId, err := g.gitlabGetProjectId(gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in getting gitlab project id")
	}
	res, err := g.gitlabGetFileRequest(projectId, "src/main/resources/application.yml", gitlabBranch)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in getting gitlab requested file")
	}

	yaml.Unmarshal(res, &yamlFile)
	helpers.Walk(yamlFile)
}

func (g *GitlabData) DeleteEnv(gitlabProject, key, env string) {
	projectId, err := g.gitlabGetProjectId(gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in getting gitlab project id")
	}

	_, err = g.gitlabVariablesDeleteRequest(projectId, key, env)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in deleting gitlab env")
	}
	fmt.Printf("The %v key in %v environment_scope is deleted\n", key, env)
}

func (g *GitlabData) AddEnvFromFile(gitlabProject, file, env string) {
	projectId, err := g.gitlabGetProjectId(gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in getting gitlab project id")
	}

	readFile, err := os.Open(file)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in opening file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		tmp := strings.Split(fileScanner.Text(), "=")
		postBody, _ := json.Marshal(map[string]string{
			"key":               tmp[0],
			"value":             tmp[1],
			"environment_scope": env,
		})

		_, err = g.gitlabVariablesPostRequest(projectId, postBody)
		if err != nil {
			helpers.CmdErrorHandler(err, "some error in posting gitlab env")
		}
		fmt.Printf("The %v key with %v value in %v environment_scope is added\n", tmp[0], tmp[1], env)

	}

	readFile.Close()
}

func (g *GitlabData) AddEnv(gitlabProject, key, value, env string) {
	projectId, err := g.gitlabGetProjectId(gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in getting gitlab project id")
	}

	postBody, _ := json.Marshal(map[string]string{
		"key":               key,
		"value":             value,
		"environment_scope": env,
	})

	_, err = g.gitlabVariablesPostRequest(projectId, postBody)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in posting gitlab env")
	}
	fmt.Printf("The %v key with %v value in %v environment_scope is added\n", key, value, env)
}

func (g *GitlabData) GenerateDockerfile(gitlabProject string) {
	projectId, err := g.gitlabGetProjectId(gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err, "some error in getting gitlab project id")
	}

	res, err := g.gitlabGetFileRequest(projectId, "pom.xml", "dev")
	if err != nil {
		panic(err)
	}

	jarFileName := helpers.GetJarFileName(res)
	fmt.Println(templates.JavaSpringDockerfile(jarFileName))
}
