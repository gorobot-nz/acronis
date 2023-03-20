package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorobot-nz/acronis/pkg/client/apimodels"
	"io"
	"net/http"
)

type fetchServices struct {
	Items []apimodels.Service `json:"items"`
}

func (c *AcronisClient) FetchServices() ([]apimodels.Service, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(applicationsUrl, c.baseUrl), nil)
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
	var fetch fetchServices
	err = json.Unmarshal(body, &fetch)
	if err != nil {
		return nil, err
	}
	return fetch.Items, nil
}

func (c *AcronisClient) ActivateService(tenantId, serviceType string) error {
	services, err := c.FetchServices()
	if err != nil {
		return err
	}

	var serviceId string
	for _, val := range services {
		if val.Type == serviceType {
			serviceId = val.Id
			break
		}
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(applicationActivateUrl, c.baseUrl, serviceId, tenantId), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}
	return nil
}
