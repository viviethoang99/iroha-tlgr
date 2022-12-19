package transferac

import (
	bizac "iroha-tlgr/appcenter/biz"
	modelac "iroha-tlgr/appcenter/model"
	serviceac "iroha-tlgr/appcenter/service"
	"iroha-tlgr/utils"
)

func GetAllAppReleaseLasted(config utils.Config, author string) string {
	androidApp := modelac.App{
		Owner:   config.ENVConfig.Owner,
		AppName: config.ENVConfig.AppNameAndroid,
	}
	iosApp := modelac.App{
		Owner:   config.ENVConfig.Owner,
		AppName: config.ENVConfig.AppNameIos,
	}
	service := serviceac.CreateAPIWithClientParams(config.AppCenterToken)
	biz := bizac.NewFindReleaseAppBiz(service)
	return biz.FindAllReleaseApp(androidApp, iosApp).TelegramReleaseMessage(author)
}
