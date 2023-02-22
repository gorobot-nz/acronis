package apimodels

const (
	AccountBillingType    = "billing"
	AccountManagementType = "management"
	AccountTechnicalType  = "technical"
)

const (
	QuotaNotification             = "quota"
	ReportsNotification           = "reports"
	BackupErrorNotification       = "backup_error"
	BackupWarningNotification     = "backup_warning"
	BackupInfoNotification        = "backup_info"
	BackupDailyReportNotification = "backup_daily_report"
)

type AccountContacts struct {
	Email     string   `json:"email"`
	Types     []string `json:"types"`
	Address1  string   `json:"address1"`
	Address2  string   `json:"address2"`
	Country   string   `json:"country"`
	State     string   `json:"state"`
	ZipCode   string   `json:"zipcode"`
	City      string   `json:"city"`
	Phone     string   `json:"phone"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastName"`
}

type Account struct {
	Id               string          `json:"id"`
	Version          uint64          `json:"version"`
	TenantId         string          `json:"tenant_id"`
	Login            string          `json:"login"`
	Contact          AccountContacts `json:"contact"`
	Activated        bool            `json:"activated"`
	Enabled          bool            `json:"enabled"`
	TermsAccepted    bool            `json:"terms_accepted"`
	CreatedAt        string          `json:"created_at"`
	Language         string          `json:"language"`
	IdpId            string          `json:"idp_id"`
	ExternalId       string          `json:"external_id"`
	PersonalTenantId string          `json:"personal_tenant_id"`
	Notifications    []string        `json:"notifications"`
	MFAStatus        string          `json:"mfa_status"`
}
