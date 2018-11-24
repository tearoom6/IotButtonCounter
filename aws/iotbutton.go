package aws

type deviceInfoAttributes struct {
	ProjectRegion      string `json:"projectRegion"`
	ProjectName        string `json:"projectName"`
	PlacementName      string `json:"placementName"`
	DeviceTemplateName string `json:"deviceTemplateName"`
}

type deviceInfo struct {
	DeviceId      string               `json:"deviceId"`
	Type          string               `json:"type"`
	RemainingLife float32              `json:"remainingLife"`
	Attributes    deviceInfoAttributes `json:"attributes"`
}

type buttonClickedEvent struct {
	ClickType    string `json:"clickType"`
	ReportedTime string `json:"reportedTime"`
}

type deviceEvent struct {
	ButtonClicked buttonClickedEvent `json:"buttonClicked"`
}

type placementInfoAttributes struct {
	DeviceId string `json:"deviceId"`
}

type placementInfoDevices map[string]string

type placementInfo struct {
	ProjectName   string                  `json:"projectName"`
	PlacementName string                  `json:"placementName"`
	Attributes    placementInfoAttributes `json:"attributes"`
	Devices       placementInfoDevices    `json:"devices"`
}

type IotEnterpriseButtonEvent struct {
	DeviceInfo    deviceInfo    `json:"deviceInfo"`
	DeviceEvent   deviceEvent   `json:"deviceEvent"`
	PlacementInfo placementInfo `json:"placementInfo"`
}
