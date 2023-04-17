package apimodels

type OfferingItemEdition string

const (
	ProtectionPerGigabyte       OfferingItemEdition = "pck_per_gigabyte"
	ProtectionPerWorkload       OfferingItemEdition = "pck_per_workload"
	FileSyncAndSharePerGigabyte OfferingItemEdition = "fss_per_gigabyte"
	FileSyncAndSharePerUser     OfferingItemEdition = "fss_per_user"
)

type OfferingItem struct {
	ApplicationId   string              `json:"application_id"`
	Edition         OfferingItemEdition `json:"edition,omitempty"`
	InfraId         string              `json:"infra_id,omitempty"`
	Locked          bool                `json:"locked"`
	MeasurementUnit string              `json:"measurement_unit"`
	Name            string              `json:"name"`
	ItemQuota       Quota               `json:"quota"`
	UsageName       string              `json:"usage_name"`
	Status          int                 `json:"status"`
	Type            string              `json:"type"`
}
