/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Gitlab Operations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(gitlabCmd)
	gitlabCmd.PersistentFlags().String("gitlab-url", "https://gitlab.com", "Gitlab url")
	gitlabCmd.PersistentFlags().String("gitlab-api-token", "", "Gitlab api token")
	gitlabCmd.PersistentFlags().StringP("project", "p", "", "Gitlab Project Name")

	gitlabCmd.MarkPersistentFlagRequired("gitlab-api-token")
	gitlabCmd.MarkPersistentFlagRequired("project")
}
