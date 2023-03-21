package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorobot-nz/acronis/pkg/client/apimodels"
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

func (c *AcronisClient) SwitchToProduction(tenantId string) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(tenantPricingUrl, c.baseUrl, tenantId), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	var tenantPricing apimodels.TenantPricing
	err = json.Unmarshal(body, &tenantPricing)
	if err != nil {
		return err
	}

	tenantPricing.Mode = apimodels.TenantProductionMode
	reqBody, err := json.Marshal(tenantPricing)
	if err != nil {
		return err
	}

	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf(tenantPricingUrl, c.baseUrl, tenantId), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err = c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}

func (c *AcronisClient) EnableItems(tenantId string, items []string) error {
	client, err := c.GetClient()
	if err != nil {
		return err
	}

	offeringItems, err := c.GetOfferingItems(client.TenantId)
	if err != nil {
		return err
	}

	var itemsMap = map[string][]apimodels.OfferingItem{
		"offering_items": offeringItems,
	}
	reqBody, err := json.Marshal(itemsMap)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(enableOfferingItemsUrl, c.baseUrl, tenantId), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}
