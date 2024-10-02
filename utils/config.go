package utils

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ProjectName     string         `json:"project_name"`
	AppcenterToken  string         `json:"appcenter_token"`
	TelegramConfig  TelegramConfig `json:"telegram_config"`
	GitlabConfig    GitlabConfig   `json:"gitlab_config"`
	ListBranchMerge []string       `json:"list_branch_merge"`
	SpecialUsers    []SpecialUser  `json:"special_users"`
	Branches        []Branch       `json:"branches"`
}

type ProjectResult struct {
	ProjectName     string         `json:"project_name"`
	AppcenterToken  string         `json:"appcenter_token"`
	TelegramConfig  TelegramConfig `json:"telegram_config"`
	GitlabConfig    GitlabConfig   `json:"gitlab_config"`
	ListBranchMerge []string       `json:"list_branch_merge"`
	SpecialUsers    []SpecialUser  `json:"special_users"`
	Branches        Branch         `json:"branches"`
}

type TelegramConfig struct {
	TelegramChatId   int    `json:"telegram_chat_id"`
	TokenBotTelegram string `json:"token_bot_telegram"`
	FileIdFailed     string `json:"file_id_failed"`
}

type GitlabConfig struct {
	AccessToken string `json:"access_token"`
	IdProject   int    `json:"id_project"`
	BaseUrl     string `json:"base_url"`
}

type SpecialUser struct {
	UserName string `json:"username"`
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type Branch struct {
	BranchName string    `json:"branch_name"`
	EnvConfig  EnvConfig `json:"env_config"`
}

type EnvConfig struct {
	AppNameAndroid string `json:"app_name_android"`
	AppNameIos     string `json:"app_name_ios"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
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

func FindConfigByNameProject(config Config, nameProject string, branchName string) (ProjectResult, error) {
	for _, project := range config.Projects {
		if project.ProjectName != nameProject {
			continue
		}
		for _, branch := range project.Branches {
			if branch.BranchName == branchName {
				return ProjectResult{
					ProjectName:     project.ProjectName,
					AppcenterToken:  project.AppcenterToken,
					TelegramConfig:  project.TelegramConfig,
					GitlabConfig:    project.GitlabConfig,
					ListBranchMerge: project.ListBranchMerge,
					SpecialUsers:    project.SpecialUsers,
					Branches:        branch,
				}, nil
			}
		}
	}

	return ProjectResult{}, errors.New("not found config by name project")
}
