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

// Args
// 1, Action
// 2, Author comm
// 3, Branch
// 4, Des
func main() {
	if len(os.Args) > 4 && strings.Contains(os.Args[4], "NO_PUSH_NOTI") {
		fmt.Println("Đóng ứng dụng")
		os.Exit(0)
	}

	branch := strings.ToLower(os.Args[3])
	urlConfig := "./config/" + branch
	config, envError := utils.LoadConfig(urlConfig)
	author := os.Args[2]

	if envError != nil {
		fmt.Println("not load config", envError)
		panic(envError)
	}

	switch os.Args[1] {
	case "build_success":
		releaseBody := transferac.GetAllAppReleaseLasted(config)
		shouldGetDataMerge := slices.Contains(config.ListBranchMerge, branch) // true
		infoMerge, _ := gitlab.GetInfoUserCreateMergeRequest(config, shouldGetDataMerge, branch)
		if infoMerge != nil && infoMerge.Author.Name != "" {
			author = infoMerge.Author.Name
		}
		content := releaseBody.TelegramReleaseMessage(author, os.Args[3])
		telegram.SentMessageToTelegram(
			config.TelegramConfig.TelegramChatId,
			content,
			config.TelegramConfig.TokenBotTelegram,
		)
	case "build_failed":
		isSpecialUser := false
		content := modelac.TelegramBuildFailed(author, config)
		for key, _ := range config.SpecialUser {
			if strings.ToLower(key) == strings.ToLower(author) {
				isSpecialUser = true
				break
			}
		}
		if !isSpecialUser {
			telegram.SentStickerToTelegram(
				config.TelegramConfig.TelegramChatId,
				config.TelegramConfig.FileIdFailed,
				config.TelegramConfig.TokenBotTelegram,
			)
		}
		telegram.SentMessageToTelegram(
			config.TelegramConfig.TelegramChatId,
			content,
			config.TelegramConfig.TokenBotTelegram,
		)
	default:
		fmt.Println("Sai tham số truyền vào: ", os.Args[1])
	}
}
