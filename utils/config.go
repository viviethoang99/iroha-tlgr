package utils

import "github.com/spf13/viper"

type Config struct {
	AppCenterToken string `mapstructure:"appcenter_token"`
	TelegramConfig struct {
		TelegramChatId   int    `mapstructure:"telegram_chat_id"`
		TokenBotTelegram string `mapstructure:"token_bot_telegram"`
	} `mapstructure:"telegram_config"`
	ENVConfig struct {
		Owner          string `mapstructure:"owner"`
		AppNameAndroid string `mapstructure:"app_name_android"`
		AppNameIos     string `mapstructure:"app_name_ios"`
	} `mapstructure:"env_config"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
