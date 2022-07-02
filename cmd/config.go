/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"alfred/operations"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set configuration for Alfred",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		data := make(map[string]string)
		gitlabUrl, _ := cmd.Flags().GetString("gitlab-url")
		gitlabApiToken, _ := cmd.Flags().GetString("gitlab-api-token")

		data["gitlab.url"] = gitlabUrl
		data["gitlab.api_token"] = gitlabApiToken

		operations.SetConfig(data)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().String("gitlab-url", "https://gitlab.com", "Set gitlab url")
	configCmd.PersistentFlags().String("gitlab-api-token", "", "Set gitlab api token")
}
