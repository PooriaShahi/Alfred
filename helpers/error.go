package helpers

import (
	"fmt"
	"os"
)

func CmdErrorHandler(err error, text string) {
	fmt.Println(text)
	fmt.Println(err)
	os.Exit(1)
}

func GitlabApiTokenNotFoundErrorHandler() {
	text := `You didn't set any gitlab api token
	Please set your gitlab api token with Config command and rerun your command`
	fmt.Println(text)
	os.Exit(1)
}

func GitlabUrlNotFoundErrorHandler() {
	text := `You didn't set any gitlab url so we use default url (https://gitlab.com)
	Or you can set your gitlab url with Config command and rerun your command`
	fmt.Println(text)
}

func UnauthorizedGitlabApiToken() {
	fmt.Println("Failed -> Your Gitlab Api Token is invalid or expired!!!")
	os.Exit(1)
}
