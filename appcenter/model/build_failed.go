package modelac

import (
	"fmt"
	"iroha-tlgr/utils"
	"strings"
)

func TelegramBuildFailed(author string, config utils.Config) string {
	for key, value := range config.SpecialUser {
		if strings.ToLower(key) == strings.ToLower(author) {
			userTag := fmt.Sprintf("[%s](tg://user?id=%s)", strings.Title(key), value)
			return userTag + " Bản build lỗi rồi kìa\\. Kiểm tra lại đi\\!"
		}
	}
	return fmt.Sprintf("Bản build của %s xảy ra lỗi rồi\\. Ai quen thì tag anh ấy vào fix nhé\\.", author)
}
