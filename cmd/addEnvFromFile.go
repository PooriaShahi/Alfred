/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/operations"

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
		gitlab := operations.GetGitlabDataObject()
		gitlabProject, _ := cmd.Flags().GetString("project")

		gitlab.AddEnvFromFile(gitlabProject, file, env)
	},
}

func init() {
	gitlabCmd.AddCommand(addEnvFromFileCmd)

	addEnvFromFileCmd.PersistentFlags().String("from-file", "", "Path of env file")
	addEnvFromFileCmd.PersistentFlags().StringP("env", "e", "*", "environment of gitlab variable")

	addEnvFromFileCmd.MarkPersistentFlagRequired("from-file")
}
