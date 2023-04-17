package apimodels

type Quota struct {
	Value   uint64 `json:"value,omitempty"`
	Overage uint64 `json:"overage,omitempty"`
	Version uint64 `json:"version"`
}
