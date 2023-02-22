package client

const (
	apiUrl             = "%s/api/2"
	tokenUrl           = "%s/idp/token"
	clientIdUrl        = "%s/clients/%s"
	tenantsUrl         = "%s/tenants"
	childTenantsUrl    = "%s/tenants/%s/children?include_details=true"
	checkLoginUrl      = "%s/users/check_login?username=%s"
	usersUrl           = "%s/users"
	fetchUser          = "%s/users/%s"
	userActivateUrl    = "%s/users/%s/send-activation-email"
	userSetPasswordUrl = "%s/users/%s/password"
	searchUrl          = "%s/search?tenant=%s&text=%s"
)
