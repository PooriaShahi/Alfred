/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"alfred/helpers"
	"alfred/operations"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// addEnvCmd represents the addEnv command
var addEnvCmd = &cobra.Command{
	Use:   "addEnv",
	Short: "Add environment variables from CLI to a specific project in gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gitlabUrl, _ := cmd.Flags().GetString("gitlab-url")
		gitlabApiToken, _ := cmd.Flags().GetString("gitlab-api-token")
		gitlabProject, _ := cmd.Flags().GetString("project")
		projectId, err := operations.GitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
		if err != nil {
			helpers.CmdErrorHandler(err)
		}

		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		env, _ := cmd.Flags().GetString("env")
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
	},
}

func init() {
	gitlabCmd.AddCommand(addEnvCmd)
	addEnvCmd.PersistentFlags().StringP("key", "k", "", "Key of env")
	addEnvCmd.PersistentFlags().StringP("value", "v", "", "Value of env")
	addEnvCmd.PersistentFlags().StringP("env", "e", "*", "environment of gitlab variable")

	addEnvCmd.MarkPersistentFlagRequired("key")
	addEnvCmd.MarkPersistentFlagRequired("value")

}
