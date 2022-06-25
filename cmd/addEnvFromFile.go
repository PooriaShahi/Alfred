/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/helpers"
	"alfred/operations"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addEnvFromFileCmd represents the addEnvFromFile command
var addEnvFromFileCmd = &cobra.Command{
	Use:   "addEnvFromFile",
	Short: "Add environment variables from a file to a specific project in gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var fileLines []string
		res := make(map[string]string)

		file, _ := cmd.Flags().GetString("from-file")
		gitlabUrl, _ := cmd.Flags().GetString("gitlab-url")
		gitlabApiToken, _ := cmd.Flags().GetString("gitlab-api-token")
		gitlabProject, _ := cmd.Flags().GetString("project")
		projectId, err := operations.GitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
		if err != nil {
			helpers.CmdErrorHandler(err)
		}

		readFile, err := os.Open(file)

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		for _, v := range fileLines {
			tmp := strings.Split(v, "=")
			res[tmp[0]] = tmp[1]
		}

		env, _ := cmd.Flags().GetString("env")
		for key, value := range res {
			postBody, _ := json.Marshal(map[string]string{
				"key":               key,
				"value":             value,
				"environment_scope": env,
			})

			_, err = operations.GitlabVariablesPostRequest(gitlabUrl, gitlabApiToken, projectId, postBody)
			if err != nil {
				helpers.CmdErrorHandler(err)
			}
			fmt.Println(fmt.Sprintf("The %v key with %v value in %v environment_scope is added", key, value, env))
		}

		readFile.Close()
	},
}

func init() {
	gitlabCmd.AddCommand(addEnvFromFileCmd)

	addEnvFromFileCmd.PersistentFlags().String("from-file", "", "Path of env file")
	addEnvFromFileCmd.PersistentFlags().StringP("env", "e", "*", "environment of gitlab variable")

	addEnvFromFileCmd.MarkPersistentFlagRequired("from-file")
}
