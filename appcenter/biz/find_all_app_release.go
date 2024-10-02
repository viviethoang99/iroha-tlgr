package bizac

import (
	"fmt"
	modelac "iroha-tlgr/appcenter/model"
)

type FindReleaseService interface {
	GetAppReleaseLastDetails(app modelac.App) (modelac.Release, error)
}

func NewFindReleaseAppBiz(service FindReleaseService) *findAllReleaseBiz {
	return &findAllReleaseBiz{service: service}
}

type findAllReleaseBiz struct {
	service FindReleaseService
}

func (biz *findAllReleaseBiz) FindAllReleaseApp(infoAndroid, infoIos modelac.App) modelac.BuildInfo {
	releaseAndroid, errAndroid := biz.service.GetAppReleaseLastDetails(infoAndroid)
	if errAndroid != nil {
		fmt.Println("không lấy được dữ liệu", errAndroid.Error())
	}
	releaseIos, errIos := biz.service.GetAppReleaseLastDetails(infoIos)
	if errIos != nil {
		fmt.Println("không lấy được dữ liệu", errIos.Error())
	}
	return modelac.BuildInfo{ApkInfo: releaseAndroid, IosInfo: releaseIos}
}
