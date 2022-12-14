package main

import (
	"fmt"
	transferac "iroha-tlgr/appcenter/transfer"
	"iroha-tlgr/telegram"
	"iroha-tlgr/utils"
	"os"
)

// Args
// 1, Action
// 2, Author commit
// 3, Decs

func main() {
	config, envError := utils.LoadConfig(".")
	if envError != nil {
		fmt.Println("not load config", envError)
		panic(envError)
	}

	switch os.Args[1] {
	case "build_success":
		releaseBody := transferac.GetAllAppReleaseLasted(config, os.Args[2], os.Args[3])
		fmt.Println(releaseBody)
		telegram.SentMessageToTelegram(
			config.TelegramConfig.TelegramChatId,
			releaseBody,
			config.TelegramConfig.TokenBotTelegram,
		)
	default:
		fmt.Println("Sai tham số truyền vào: ", os.Args[1])
	}
}
