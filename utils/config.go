package utils

import "github.com/spf13/viper"

type Config struct {
	NameBranch     string `mapstructure:"name_branch"`
	AppCenterToken string `mapstructure:"appcenter_token"`
	TelegramConfig struct {
		TelegramChatId   int    `mapstructure:"telegram_chat_id"`
		TokenBotTelegram string `mapstructure:"token_bot_telegram"`
		FileIdFailed     string `mapstructure:"file_id_failed"`
	} `mapstructure:"telegram_config"`
	ENVConfig struct {
		Owner          string `mapstructure:"owner"`
		AppNameAndroid string `mapstructure:"app_name_android"`
		AppNameIos     string `mapstructure:"app_name_ios"`
	} `mapstructure:"env_config"`
	GitlabConfig struct {
		AccessToken string `mapstructure:"access_token"`
		IdProject   int    `mapstructure:"id_project"`
		BaseUrl     string `mapstructure:"base_url"`
	} `mapstructure:"gitlab_config"`
	SpecialUsers    []SpecialUser `mapstructure:"special_users"`
	ListBranchMerge []string      `mapstructure:"list_branch_merge"`
}

type SpecialUser struct {
	UserName string `mapstructure:"username"`
	ID       string `mapstructure:"id"`
	FullName string `mapstructure:"full_name"`
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
