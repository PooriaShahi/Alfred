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

	"gopkg.in/yaml.v3"
)

var client http.Client

func gitlabGetProjectId(gitlabUrl string, gitlabApiToken string, gitlabProject string) (string, error) {
	var data []map[string]interface{}
	var projectId string
	Url := gitlabUrl + "/api/v4/search?scope=projects&search=" + gitlabProject
	req, err := http.NewRequest("GET", Url, nil)
	req.Header.Set("PRIVATE-TOKEN", gitlabApiToken)

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if res.StatusCode == 401 {
		fmt.Println("Failed -> Your Gitlab Api Token is invalid or expired!!!")
		os.Exit(1)
	}

	result, err := ioutil.ReadAll(res.Body)
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

func gitlabVariablesPostRequest(gitlabUrl string, gitlabApiToken string, projectId string, postBody []byte) ([]byte, error) {
	url := gitlabUrl + "/api/v4/projects/" + projectId + "/variables"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	req.Header.Set("PRIVATE-TOKEN", gitlabApiToken)
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

func gitlabVariablesDeleteRequest(gitlabUrl string, gitlabApiToken string, projectId string, key string, env string) ([]byte, error) {
	url := gitlabUrl + "/api/v4/projects/" + projectId + "/variables/" + key + "?filter[environment_scope]=" + env
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("PRIVATE-TOKEN", gitlabApiToken)

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

func gitlabGetFileRequest(gitlabUrl string, gitlabApiToken string, projectId string, filePath string, branch string) ([]byte, error) {
	var resultData map[string]interface{}
	filePath = strings.ReplaceAll(filePath, "/", "%2F")
	filePath = strings.ReplaceAll(filePath, ".", "%2E")
	url := gitlabUrl + "/api/v4/projects/" + projectId + "/repository/files/" + filePath + "?ref=" + branch
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", gitlabApiToken)

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

func GenerateEnvsSpringBoot(gitlabUrl, gitlabApiToken, gitlabProject, gitlabBranch string) {
	var yamlFile map[string]interface{}

	projectId, err := gitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}
	res, err := gitlabGetFileRequest(gitlabUrl, gitlabApiToken, projectId, "src/main/resources/application.yml", gitlabBranch)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}

	yaml.Unmarshal(res, &yamlFile)
	helpers.Walk(yamlFile)
}

func DeleteEnv(gitlabUrl, gitlabApiToken, gitlabProject, key, env string) {
	projectId, err := gitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}

	_, err = gitlabVariablesDeleteRequest(gitlabUrl, gitlabApiToken, projectId, key, env)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}
	fmt.Println(fmt.Sprintf("The %v key in %v environment_scope is deleted", key, env))
}

func AddEnvFromFile(gitlabUrl, gitlabApiToken, gitlabProject, file, env string) {
	projectId, err := gitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}

	readFile, err := os.Open(file)
	if err != nil {
		helpers.CmdErrorHandler(err)
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

		_, err = gitlabVariablesPostRequest(gitlabUrl, gitlabApiToken, projectId, postBody)
		if err != nil {
			helpers.CmdErrorHandler(err)
		}
		fmt.Println(fmt.Sprintf("The %v key with %v value in %v environment_scope is added", tmp[0], tmp[1], env))

	}

	readFile.Close()
}

func AddEnv(gitlabUrl, gitlabApiToken, gitlabProject, key, value, env string) {
	projectId, err := gitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}

	postBody, _ := json.Marshal(map[string]string{
		"key":               key,
		"value":             value,
		"environment_scope": env,
	})

	_, err = gitlabVariablesPostRequest(gitlabUrl, gitlabApiToken, projectId, postBody)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}
	fmt.Println(fmt.Sprintf("The %v key with %v value in %v environment_scope is added", key, value, env))
}

func GenerateDockerfile(gitlabUrl, gitlabApiToken, gitlabProject string) {
	projectId, err := gitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
	if err != nil {
		helpers.CmdErrorHandler(err)
	}

	res, err := gitlabGetFileRequest(gitlabUrl, gitlabApiToken, projectId, "pom.xml", "dev")
	if err != nil {
		panic(err)
	}

	jarFileName := helpers.GetJarFileName(res)
	fmt.Println(templates.JavaSpringDockerfile(jarFileName))
}
