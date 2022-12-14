package telegram

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func SentMessageToTelegram(chatId int, text string, token string) {
	var telegramApi = "https://api.telegram.org/bot" + token + "/sendMessage"
	parseMode := "MarkdownV2"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id":    {strconv.Itoa(chatId)},
			"parse_mode": {parseMode},
			"text":       {text},
		})
	if err != nil {
		log.Fatalln("không gửi được tin nhắn: ", err.Error())
	}

	defer response.Body.Close()

	var bodyType, errRead = io.ReadAll(response.Body)
	if errRead != nil {
		log.Fatalln("Không parse được dữ liệu: ", err.Error())
	}
	bodyString := string(bodyType)
	log.Println("Body of Telegram Response:", bodyString)
}
