package apimodels

const (
	TenantPartnerKind  = "partner"
	TenantFolderKind   = "folder"
	TenantCustomerKind = "customer"
	TenantUnitKind     = "unit"
)

type TenantContacts struct {
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	State     string `json:"state"`
	ZipCode   string `json:"zipcode"`
}

type TenantCreate struct {
	Name        string         `json:"name"`
	Kind        string         `json:"kind"`
	ParentId    string         `json:"parent_id"`
	InternalTag string         `json:"internal_tag"`
	Language    string         `json:"language"`
	Contact     TenantContacts `json:"contact"`
}

type Tenant struct {
	Id              string                 `json:"id"`
	AncestralAccess bool                   `json:"ancestral_access"`
	BrandId         uint64                 `json:"brand_id"`
	BrandUUID       string                 `json:"brand_uuid"`
	Contact         TenantContacts         `json:"contact"`
	CustomerId      string                 `json:"customer_id"`
	CustomerType    string                 `json:"customer_type"`
	DefaultIdpId    string                 `json:"default_idp_id"`
	Enabled         bool                   `json:"enabled"`
	HasChildren     bool                   `json:"has_children"`
	InternalTag     string                 `json:"internal_tag"`
	Kind            string                 `json:"kind"`
	Language        string                 `json:"language"`
	Name            string                 `json:"name"`
	OwnerId         string                 `json:"owner_id"`
	ParentId        string                 `json:"parent_id"`
	UpdateLock      map[string]interface{} `json:"update_lock"`
	Version         uint64                 `json:"version"`
}
