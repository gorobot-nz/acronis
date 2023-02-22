package apimodels

const (
	TenantPartnerKind  = "partner"
	TenantFolderKind   = "folder"
	TenantCustomerKind = "customer"
	TenantUnitKind     = "unit"
)

type TenantContacts struct {
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Phone     string `json:"phone,omitempty"`
	State     string `json:"state,omitempty"`
	ZipCode   string `json:"zipcode,omitempty"`
}

type Tenant struct {
	Id              string                 `json:"id,omitempty"`
	AncestralAccess bool                   `json:"ancestral_access,omitempty"`
	BrandId         uint64                 `json:"brand_id,omitempty"`
	BrandUUID       string                 `json:"brand_uuid,omitempty"`
	Contact         TenantContacts         `json:"contact,omitempty"`
	CustomerId      string                 `json:"customer_id,omitempty"`
	CustomerType    string                 `json:"customer_type,omitempty"`
	DefaultIdpId    string                 `json:"default_idp_id,omitempty"`
	Enabled         bool                   `json:"enabled,omitempty"`
	HasChildren     bool                   `json:"has_children,omitempty"`
	InternalTag     string                 `json:"internal_tag,omitempty"`
	Kind            string                 `json:"kind,omitempty"`
	Language        string                 `json:"language,omitempty"`
	Name            string                 `json:"name,omitempty"`
	OwnerId         string                 `json:"owner_id,omitempty"`
	ParentId        string                 `json:"parent_id,omitempty"`
	UpdateLock      map[string]interface{} `json:"update_lock,omitempty"`
	Version         uint64                 `json:"version,omitempty"`
}
