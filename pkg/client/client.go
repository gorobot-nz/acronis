package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	apimodels2 "github.com/gorobot-nz/acronis/pkg/client/apimodels"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	baseUrl := fmt.Sprintf(apiUrl, datacenterUrl)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))
	data := url.Values{}
	data.Add(grantType, credentialsGrantType)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(tokenUrl, baseUrl), strings.NewReader(data.Encode()))
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
	var authResponse = apimodels2.AuthResponse{}
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

func (c *AcronisClient) GetClient() (*apimodels2.Client, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(clientIdUrl, c.baseUrl, c.clientId), nil)
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

	var client = apimodels2.Client{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
