package apimodels

const (
	UserBillingType    = "billing"
	UserManagementType = "management"
	UserTechnicalType  = "technical"
)

const (
	QuotaNotification             = "quota"
	ReportsNotification           = "reports"
	BackupErrorNotification       = "backup_error"
	BackupWarningNotification     = "backup_warning"
	BackupInfoNotification        = "backup_info"
	BackupDailyReportNotification = "backup_daily_report"
)

type UserContacts struct {
	Email     string   `json:"email,omitempty"`
	Types     []string `json:"types,omitempty"`
	Address1  string   `json:"address1,omitempty"`
	Address2  string   `json:"address2,omitempty"`
	Country   string   `json:"country,omitempty"`
	State     string   `json:"state,omitempty"`
	ZipCode   string   `json:"zipcode,omitempty"`
	City      string   `json:"city,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
}

type User struct {
	Id               string       `json:"id,omitempty"`
	Version          uint64       `json:"version,omitempty"`
	TenantId         string       `json:"tenant_id,omitempty"`
	Login            string       `json:"login,omitempty"`
	Contact          UserContacts `json:"contact,omitempty"`
	Activated        bool         `json:"activated,omitempty"`
	Enabled          bool         `json:"enabled,omitempty"`
	TermsAccepted    bool         `json:"terms_accepted,omitempty"`
	CreatedAt        string       `json:"created_at,omitempty"`
	Language         string       `json:"language,omitempty"`
	IdpId            string       `json:"idp_id,omitempty"`
	ExternalId       string       `json:"external_id,omitempty"`
	PersonalTenantId string       `json:"personal_tenant_id,omitempty"`
	Notifications    []string     `json:"notifications,omitempty"`
	MFAStatus        string       `json:"mfa_status,omitempty"`
}
