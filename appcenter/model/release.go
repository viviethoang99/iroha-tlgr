package modelac

import "time"

// Release ...
type Release struct {
	ID                 int       `json:"id"`
	AppName            string    `json:"app_name"`
	AppDisplayName     string    `json:"app_display_name"`
	AppOs              string    `json:"app_os"`
	Version            string    `json:"version"`
	Origin             string    `json:"origin"`
	ShortVersion       string    `json:"short_version"`
	ReleaseNotes       string    `json:"release_notes"`
	Size               int       `json:"size"`
	MinOs              string    `json:"min_os"`
	DeviceFamily       string    `json:"device_family"`
	AndroidMinAPILevel string    `json:"android_min_api_level"`
	BundleIdentifier   string    `json:"bundle_identifier"`
	PackageHashes      []string  `json:"package_hashes"`
	Fingerprint        string    `json:"fingerprint"`
	UploadedAt         time.Time `json:"uploaded_at"`
	DownloadURL        string    `json:"download_url"`
	AppIconURL         string    `json:"app_icon_url"`
	InstallURL         string    `json:"install_url"`
	DestinationType    string    `json:"destination_type"`
	Enabled            bool      `json:"enabled"`
	Status             string    `json:"status"`
	IsExternalBuild    bool      `json:"is_external_build"`
	DistributionGroups []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"distribution_groups"`
	Owner struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Error Error `json:"error"`
}
