package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	apimodels "github.com/gorobot-nz/acronis/pkg/client/apimodels"
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
	targetUrl    string
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
		targetUrl:    datacenterUrl,
	}, nil
}

func (c *AcronisClient) GetClient() (*apimodels.Client, error) {
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

	var client = apimodels.Client{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *AcronisClient) ExternalLogin(userId string) (string, error) {
	reqBody := map[string]string{
		"user_id": userId,
		"purpose": "user_login",
	}

	marshal, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(externalTokenUrl, c.baseUrl), bytes.NewBuffer(marshal))
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

	response := map[string]string{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	ott, ok := response["ott"]

	if !ok {
		return "", errors.New("no ott")
	}

	ott = url.QueryEscape(ott)

	return fmt.Sprintf(redirectUrl, c.baseUrl, ott, c.targetUrl), nil
}
