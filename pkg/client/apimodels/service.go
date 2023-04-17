package apimodels

type ServiceType string

const (
	BackupType               ServiceType = "backup"
	FileSyncAndShareType     ServiceType = "files_cloud"
	PhysicalDataShippingType ServiceType = "physical_data_shipping"
	NotaryType               ServiceType = "notary"
)

type Service struct {
	Id         string      `json:"id"`
	ApiBaseUrl string      `json:"api_base_url"`
	Name       string      `json:"name"`
	Type       ServiceType `json:"type"`
	Usages     []string    `json:"usages"`
}
