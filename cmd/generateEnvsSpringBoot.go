/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"alfred/operations"

	"github.com/spf13/cobra"
)

// generateEnvsSpringBootCmd represents the generateEnvsSpringBoot command
var generateEnvsSpringBootCmd = &cobra.Command{
	Use:   "generateEnvsSpringBoot",
	Short: "Generate Environment Variables for Java Spring Boot Project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gitlab := operations.GetGitlabDataObject()
		gitlabProject, _ := cmd.Flags().GetString("project")
		gitlabBranch, _ := cmd.Flags().GetString("branch")

		gitlab.GenerateEnvsSpringBoot(gitlabProject, gitlabBranch)
	},
}

func init() {
	gitlabCmd.AddCommand(generateEnvsSpringBootCmd)

	generateEnvsSpringBootCmd.PersistentFlags().StringP("branch", "b", "", "branch of gitlab repository")

	generateEnvsSpringBootCmd.MarkPersistentFlagRequired("branch")
}
