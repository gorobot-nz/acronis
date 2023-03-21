package apimodels

type Role struct {
	TenantId    string `json:"tenant_id"`
	TrusteeId   string `json:"trustee_id"`
	RoleId      string `json:"role_id"`
	Id          string `json:"id"`
	IssuerId    string `json:"issuer_id"`
	TrusteeType string `json:"trustee_type"`
	Version     int    `json:"version"`
}
