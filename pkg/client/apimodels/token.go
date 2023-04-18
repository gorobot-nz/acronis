package apimodels

type Token struct {
	Id         int      `json:"id"`
	Token      string   `json:"token"`
	TenantId   string   `json:"tenant_id"`
	TenantUuid string   `json:"tenant_uuid"`
	CreatedAt  string   `json:"created_at"`
	ExpiresAt  string   `json:"expires_at"`
	Scopes     []string `json:"scopes"`
}
