/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"alfred/helpers"
	"alfred/operations"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteEnvCmd represents the deleteEnv command
var deleteEnvCmd = &cobra.Command{
	Use:   "deleteEnv",
	Short: "Delete environment variables from CLI or from a file to a specific project in gitlab",
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
		env, _ := cmd.Flags().GetString("env")

		_, err = operations.GitlabVariablesDeleteRequest(gitlabUrl, gitlabApiToken, projectId, key, env)
		if err != nil {
			helpers.CmdErrorHandler(err)
		}
		fmt.Println(fmt.Sprintf("The %v key in %v environment_scope is deleted", key, env))
	},
}

func init() {
	gitlabCmd.AddCommand(deleteEnvCmd)

	deleteEnvCmd.PersistentFlags().String("from-file", "false", "Path of env file")
	deleteEnvCmd.PersistentFlags().StringP("key", "k", "", "Key of env")
	deleteEnvCmd.PersistentFlags().StringP("env", "e", "*", "environment of gitlab variable")

	deleteEnvCmd.MarkPersistentFlagRequired("key")
}
