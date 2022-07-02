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
		gitlab := operations.GetGitlabDataObject()
		gitlabProject, _ := cmd.Flags().GetString("project")

		gitlab.GenerateDockerfile(gitlabProject)
	},
}

func init() {
	gitlabCmd.AddCommand(generateDockerfileCmd)

	generateDockerfileCmd.PersistentFlags().String("foo", "", "A help for foo")

}
