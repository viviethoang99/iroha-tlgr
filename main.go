package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	modelac "iroha-tlgr/appcenter/model"
	transferac "iroha-tlgr/appcenter/transfer"
	"iroha-tlgr/gitlab"
	"iroha-tlgr/telegram"
	"iroha-tlgr/utils"
	"os"
	"strings"
)

// Constants
const (
	ActionBuildSuccess = "build_success"
	ActionBuildFailed  = "build_failed"
	ErrorLoadConfig    = "not load config"
	ErrorFindConfig    = "not find config by name project"
	ErrorInvalidParam  = "Sai tham số truyền vào: "
)

// Args
// 1, Action
// 2, Author comm
// 3, Branch
// 4, Name project
func main() {
	branch := os.Args[3]
	config, err := utils.LoadConfig()
	if err != nil {
		handleError(ErrorLoadConfig, err)
	}

	author := os.Args[2]
	configProject, err := utils.FindConfigByNameProject(config, os.Args[4], branch)
	if err != nil {
		handleError(ErrorFindConfig, err)
	}

	switch os.Args[1] {
	case ActionBuildSuccess:
		handleBuildSuccess(config, configProject, branch, author)
	case ActionBuildFailed:
		handleBuildFailed(configProject, author)
	default:
		fmt.Println(ErrorInvalidParam, os.Args[1])
	}
}

func handleError(message string, err error) {
	fmt.Println(message, err)
	panic(err)
}

func handleBuildSuccess(config utils.Config, configProject utils.ProjectResult, branch, author string) {
	releaseBody := transferac.GetAllAppReleaseLasted(config)
	shouldGetDataMerge := slices.Contains(configProject.ListBranchMerge, branch)
	infoMerge, _ := gitlab.GetInfoUserCreateMergeRequest(config, shouldGetDataMerge, branch)
	if infoMerge != nil && infoMerge.Author.Name != "" {
		author = infoMerge.Author.Name
	}
	content := releaseBody.TelegramReleaseMessage(author, configProject.Branches.BranchName)
	telegram.SentMessageToTelegram(
		configProject.TelegramConfig.TelegramChatId,
		content,
		configProject.TelegramConfig.TokenBotTelegram,
	)
}

func handleBuildFailed(configProject utils.ProjectResult, author string) {
	isSpecialUser := isSpecialUser(configProject.SpecialUsers, author)
	content := modelac.TelegramBuildFailed(author, configProject)
	if !isSpecialUser {
		telegram.SentStickerToTelegram(
			configProject.TelegramConfig.TelegramChatId,
			configProject.TelegramConfig.FileIdFailed,
			configProject.TelegramConfig.TokenBotTelegram,
		)
	}
	telegram.SentMessageToTelegram(
		configProject.TelegramConfig.TelegramChatId,
		content,
		configProject.TelegramConfig.TokenBotTelegram,
	)
}

func isSpecialUser(specialUsers []utils.SpecialUser, author string) bool {
	for _, spUser := range specialUsers {
		if strings.EqualFold(spUser.UserName, author) {
			return true
		}
	}
	return false
}
