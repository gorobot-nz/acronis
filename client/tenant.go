package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorobot-nz/acronis/client/apimodels"
	"io"
	"net/http"
)

func (c *AcronisClient) CreateTenant(tenantCreation *apimodels.Tenant) (*apimodels.Tenant, error) {
	reqBody, err := json.Marshal(*tenantCreation)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(tenantsUrl, c.baseUrl), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tenant apimodels.Tenant
	err = json.Unmarshal(body, &tenant)
	if err != nil {
		return nil, err
	}

	return &tenant, nil
}

func (c *AcronisClient) FetchTenants() string {
	client, err := c.GetClient()
	if err != nil {
		return ""
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(childTenantsUrl, c.baseUrl, client.TenantId), nil)
	if err != nil {
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(all)
}