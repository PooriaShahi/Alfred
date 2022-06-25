package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

var client http.Client

func GitlabGetProjectId(gitlabUrl string, gitlabApiToken string, gitlabProject string) (string, error) {
	var data []map[string]interface{}
	var projectId string
	Url := gitlabUrl + "/api/v4/search?scope=projects&search=" + gitlabProject
	req, err := http.NewRequest("GET", Url, nil)
	req.Header.Set("PRIVATE-TOKEN", gitlabApiToken)

	res, err := client.Do(req)
	if err != nil {
		return "", err
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

func GitlabVariablesPostRequest(gitlabUrl string, gitlabApiToken string, projectId string, postBody []byte) ([]byte, error) {
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

func GitlabVariablesDeleteRequest(gitlabUrl string, gitlabApiToken string, projectId string, key string, env string) ([]byte, error) {
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

func GitlabGetFileRequest(gitlabUrl string, gitlabApiToken string, projectId string, filename string, extension string, branch string) ([]byte, error) {
	url := gitlabUrl + "/api/v4/projects/" + projectId + "/repository/files/src%2Fmain%2Fresources%2F" + filename + "%2E" + extension + "?ref=" + branch
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
	return result, nil
}

func Walk(m map[string]interface{}) {
	for k, v := range m {

		kind := reflect.ValueOf(v).Kind()

		if kind == reflect.Map {
			new := make(map[string]interface{})
			newValue := v.(map[string]interface{})
			for key, value := range newValue {
				new[k+"_"+key] = value
			}
			Walk(new)
		}

		if kind == reflect.Int {
			fmt.Printf("%v=%v\n", strings.ToUpper(k), v)
		}

		if kind == reflect.String {
			fmt.Printf("%v=%v\n", strings.ToUpper(k), v)
		}

		if kind == reflect.Bool {
			fmt.Printf("%v=%v\n", strings.ToUpper(k), v)
		}
	}
}
