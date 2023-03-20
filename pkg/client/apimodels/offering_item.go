package apimodels

type OfferingItem struct {
	ApplicationId   string `json:"application_id"`
	Edition         string `json:"edition"`
	InfraId         string `json:"infra_id"`
	Locked          bool   `json:"locked"`
	MeasurementUnit string `json:"measurement_unit"`
	Name            string `json:"name"`
	ItemQuota       Quota  `json:"quota"`
	UsageName       string `json:"usage_name"`
	Status          int    `json:"status"`
	Type            string `json:"type"`
}
