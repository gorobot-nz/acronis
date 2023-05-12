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

type fetchItems struct {
	Items []apimodels.OfferingItem
}

func (c *AcronisClient) FetchOfferingItemsForChild(withEdition bool) ([]apimodels.OfferingItem, error) {
	client, err := c.GetClient()
	if err != nil {
		return nil, err
	}

	var url string
	if withEdition {
		url = getOfferingItemsForChildWithEdition
	} else {
		url = getOfferingItemsForChild
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(url, c.baseUrl, client.TenantId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var fetch fetchItems
	err = json.Unmarshal(body, &fetch)
	if err != nil {
		return nil, err
	}
	return fetch.Items, nil
}

func (c *AcronisClient) EnableOfferingItems(tenantId string, items []apimodels.OfferingItem) error {
	var body = map[string][]apimodels.OfferingItem{
		"offering_items": items,
	}

	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(enableOfferingItemsUrl, c.baseUrl, tenantId), bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}

func (c *AcronisClient) EnableOfferingItem(tenantId string, item *apimodels.OfferingItem) error {
	var body = map[string][]apimodels.OfferingItem{
		"offering_items": {*item},
	}

	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(enableOfferingItemsUrl, c.baseUrl, tenantId), bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}
	return nil
}
