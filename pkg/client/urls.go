package client

const (
	apiUrl                 = "%s/api/2"
	tokenUrl               = "%s/idp/token"
	clientIdUrl            = "%s/clients/%s"
	tenantsUrl             = "%s/tenants"
	enableOfferingItemsUrl = "%s/tenants/%s/offering_items"
	getOfferingItems       = "%s/tenants/%s/offering_items/available_for_child?kind=customer"
	childTenantsUrl        = "%s/tenants/%s/children?include_details=true"
	tenantPricingUrl       = "%s/tenants/%s/pricing"
	checkLoginUrl          = "%s/users/check_login?username=%s"
	usersUrl               = "%s/users"
	fetchUser              = "%s/users/%s"
	userActivateUrl        = "%s/users/%s/send-activation-email"
	userSetPasswordUrl     = "%s/users/%s/password"
	searchUrl              = "%s/search?tenant=%s&text=%s"
	applicationsUrl        = "%s/applications"
	applicationActivateUrl = "%s/applications/%s/bindings/tenants/%s"
)
