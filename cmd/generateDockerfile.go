/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/operations"

	"github.com/spf13/cobra"
)

// generateDockerfileCmd represents the generateDockerfile command
var generateDockerfileCmd = &cobra.Command{
	Use:   "generateDockerfile",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gitlabUrl, _ := cmd.Flags().GetString("gitlab-url")
		gitlabApiToken, _ := cmd.Flags().GetString("gitlab-api-token")
		gitlabProject, _ := cmd.Flags().GetString("project")

		operations.GenerateDockerfile(gitlabUrl, gitlabApiToken, gitlabProject)
	},
}

func init() {
	gitlabCmd.AddCommand(generateDockerfileCmd)

	generateDockerfileCmd.PersistentFlags().String("foo", "", "A help for foo")

}
