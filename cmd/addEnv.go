/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"alfred/operations"

	"github.com/spf13/cobra"
)

// addEnvCmd represents the addEnv command
var addEnvCmd = &cobra.Command{
	Use:   "addEnv",
	Short: "Add environment variables from CLI to a specific project in gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gitlab := operations.GetGitlabDataObject()
		gitlabProject, _ := cmd.Flags().GetString("project")
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		env, _ := cmd.Flags().GetString("env")

		gitlab.AddEnv(gitlabProject, key, value, env)
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
