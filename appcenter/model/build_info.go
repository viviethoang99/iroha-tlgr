package modelac

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL = `https://install.appcenter.ms`
)

type BuildInfo struct {
	ApkInfo Release
	IosInfo Release
}

func (buildInfo BuildInfo) TelegramReleaseMessage(author string, branch string) string {
	var resultString string
	apkSize := buildInfo.ApkInfo.Size >> 20
	iosSize := buildInfo.IosInfo.Size >> 20
	updateAt := time.Now().Format("15h04 - 02/01/2006 ")
	resultString += `✨ *Cập nhật lúc ` + updateAt + `* ✨` + "\n"
	resultString += "------------------------- \n"
	if branch != "" {
		resultString += `*Nhánh:* #` + branch + "\n"
	}
	if author != "" {
		resultString += `*Người build:* ` + author + "\n"
	}
	if buildInfo.IosInfo.ShortVersion != "" && buildInfo.IosInfo.Version != "" {
		resultString += "*Version:* " + buildInfo.IosInfo.ShortVersion + ` - ` + buildInfo.IosInfo.Version + "\n"
	}
	if !(apkSize == 0 && iosSize == 0) {
		resultString += "*Kích thước tệp:* \n"
	}
	if apkSize != 0 {
		resultString += " *- APK:* " + strconv.Itoa(apkSize) + "MB\n"
	}
	if iosSize != 0 {
		resultString += " *- iOS:* " + strconv.Itoa(iosSize) + "MB\n"
	}
	if !(buildInfo.ApkInfo.DownloadURL == "" && buildInfo.IosInfo.DownloadURL == "") {
		resultString += "*Link vào App Center:* \n"
	}
	if buildInfo.ApkInfo.ID != 0 {
		var releaseShowURL = fmt.Sprintf("[%s](%s/users/%s/apps/%s/releases/%d)",
			buildInfo.ApkInfo.UploadedAt.Local().Format("15h04 - 02/01/2006"),
			baseURL,
			buildInfo.ApkInfo.Owner.Name,
			buildInfo.ApkInfo.AppName,
			buildInfo.ApkInfo.ID,
		)
		resultString += " *- Android:* " + releaseShowURL + "\n"
	}
	if buildInfo.IosInfo.ID != 0 {
		var releaseShowURL = fmt.Sprintf("[%s](%s/users/%s/apps/%s/releases/%d)",
			buildInfo.IosInfo.UploadedAt.Local().Format("15h04 - 02/01/2006"),
			baseURL,
			buildInfo.IosInfo.Owner.Name,
			buildInfo.IosInfo.AppName,
			buildInfo.IosInfo.ID,
		)
		resultString += " *- iOS:* " + releaseShowURL + "\n"
	}
	return toEscapeMsg(resultString)
}

func toEscapeMsg(content string) string {
	r := strings.NewReplacer(
		"-", "\\-",
		".", "\\.",
		"=", "\\=",
		"_", "\\_",
		"#", "\\#",
	)
	return r.Replace(content)
}
