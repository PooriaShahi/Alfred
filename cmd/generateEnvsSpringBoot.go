/*
Copyright © 2022 Pooria Shahi <PooriaPro@gmail.com>

*/
package cmd

import (
	"alfred/helpers"
	"alfred/operations"
	"encoding/base64"
	"encoding/json"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// generateEnvsSpringBootCmd represents the generateEnvsSpringBoot command
var generateEnvsSpringBootCmd = &cobra.Command{
	Use:   "generateEnvsSpringBoot",
	Short: "Generate Environment Variables for Java Spring Boot Project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var resultData map[string]interface{}
		var yamlFile map[string]interface{}
		gitlabUrl, _ := cmd.Flags().GetString("gitlab-url")
		gitlabApiToken, _ := cmd.Flags().GetString("gitlab-api-token")
		gitlabProject, _ := cmd.Flags().GetString("project")
		gitlabBranch, _ := cmd.Flags().GetString("branch")
		projectId, err := operations.GitlabGetProjectId(gitlabUrl, gitlabApiToken, gitlabProject)
		if err != nil {
			helpers.CmdErrorHandler(err)
		}
		res, err := operations.GitlabGetFileRequest(gitlabUrl, gitlabApiToken, projectId, "application", "yml", gitlabBranch)
		if err != nil {
			helpers.CmdErrorHandler(err)
		}
		json.Unmarshal(res, &resultData)
		decodedData, _ := base64.StdEncoding.DecodeString(resultData["content"].(string))
		yaml.Unmarshal(decodedData, &yamlFile)
		operations.Walk(yamlFile)
	},
}

func init() {
	gitlabCmd.AddCommand(generateEnvsSpringBootCmd)

	generateEnvsSpringBootCmd.PersistentFlags().StringP("branch", "b", "", "branch of gitlab repository")

	generateEnvsSpringBootCmd.MarkPersistentFlagRequired("branch")
}
