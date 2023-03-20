package apimodels

type Service struct {
	Id         string   `json:"id"`
	ApiBaseUrl string   `json:"api_base_url"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Usages     []string `json:"usages"`
}
