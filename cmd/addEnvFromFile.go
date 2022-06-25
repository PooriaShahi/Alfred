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
		file, _ := cmd.Flags().GetString("from-file")
		env, _ := cmd.Flags().GetString("env")
		gitlabUrl, _ := cmd.Flags().GetString("gitlab-url")
		gitlabApiToken, _ := cmd.Flags().GetString("gitlab-api-token")
		gitlabProject, _ := cmd.Flags().GetString("project")
		projectId, err := operations.GitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
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

			_, err = operations.GitlabVariablesPostRequest(gitlabUrl, gitlabApiToken, projectId, postBody)
			if err != nil {
				helpers.CmdErrorHandler(err)
			}
			fmt.Println(fmt.Sprintf("The %v key with %v value in %v environment_scope is added", tmp[0], tmp[1], env))

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
