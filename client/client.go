package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorobot-nz/acronis/client/apimodels"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorobot-nz/acronis/client/consts"
	"github.com/gorobot-nz/acronis/client/urls"
)

type AcronisClient struct {
	*http.Client
	clientId     string
	clientSecret string
	baseUrl      string
	token        string
}

func NewAcronisClient(clientId, clientSecret, datacenterUrl string) (*AcronisClient, error) {
	httpClient := &http.Client{
		Transport: nil,
	}
	baseUrl := fmt.Sprintf(urls.ApiUrl, datacenterUrl)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))
	data := url.Values{}
	data.Add(consts.GrantType, consts.CredentialsGrantType)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(urls.TokenUrl, baseUrl), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCredentials))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authResponse = apimodels.AuthResponse{}
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}

	return &AcronisClient{
		Client:       httpClient,
		clientId:     clientId,
		clientSecret: clientSecret,
		baseUrl:      baseUrl,
		token:        authResponse.AccessToken,
	}, nil
}

func (c *AcronisClient) GetClient() (*apimodels.Client, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(urls.ClientIdUrl, c.baseUrl, c.clientId), nil)
	if err != nil {
		return nil, err
	}
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

	var client = apimodels.Client{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *AcronisClient) CreateCustomerTenant(name string) (string, error) {
	client, err := c.GetClient()
	if err != nil {
		return "", err
	}

	var customer = map[string]string{
		"name":      name,
		"kind":      "customer",
		"parent_id": client.TenantId,
	}

	reqBody, err := json.Marshal(customer)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(urls.TenantsUrl, c.baseUrl), bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tenant apimodels.Tenant
	err = json.Unmarshal(body, &tenant)
	if err != nil {
		return "", err
	}

	return tenant.Id, nil
}
