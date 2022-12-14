package serviceac

import (
	"fmt"
	"iroha-tlgr/appcenter/model"
	"net/http"
)

func (api *api) GetAppReleaseLastDetails(app modelac.App) (modelac.Release, error) {
	var (
		releaseShowURL = fmt.Sprintf("%s/v0.1/apps/%s/%s/releases/latest?is_install_page=true",
			baseURL,
			app.Owner,
			app.AppName,
		)
		release modelac.Release
	)

	statusCode, err := api.client.jsonRequest(http.MethodGet, releaseShowURL, nil, &release)
	if err != nil {
		return modelac.Release{}, err
	}

	if statusCode != http.StatusOK {
		return modelac.Release{}, fmt.Errorf(
			"invalid status code: %d, url: %s, body: %v",
			statusCode,
			releaseShowURL,
			release,
		)
	}

	return release, err
}
