package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	apimodels "github.com/gorobot-nz/acronis/pkg/client/apimodels"
	"io"
	"net/http"
)

func (c *AcronisClient) checkLogin(login string) bool {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(checkLoginUrl, c.baseUrl, login), nil)
	if err != nil {
		return false
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.Do(req)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusNoContent {
		return false
	}
	return true
}

func (c *AcronisClient) CreateUser(userCreate *apimodels.User) (*apimodels.User, error) {
	if isLogin := c.checkLogin(userCreate.Login); !isLogin {
		return nil, errors.New("login is taken")
	}

	reqBody, err := json.Marshal(*userCreate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(usersUrl, c.baseUrl), bytes.NewBuffer(reqBody))
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

	var user apimodels.User
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type activateRoles struct {
	Items []apimodels.Role `json:"items"`
}

func (c *AcronisClient) ActivateRoles(userId, tenantId string) string {
	var roles = activateRoles{Items: []apimodels.Role{
		apimodels.Role{
			TenantId:    tenantId,
			TrusteeId:   userId,
			RoleId:      "backup_user",
			Id:          "00000000-0000-0000-0000-000000000000",
			IssuerId:    "00000000-0000-0000-0000-000000000000",
			TrusteeType: "user",
			Version:     0,
		},
	}}

	marshal, err := json.Marshal(roles)
	if err != nil {
		return ""
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(usersUrl, c.baseUrl), bytes.NewBuffer(marshal))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

func (c *AcronisClient) ActivateWithPassword(userId, password string) error {
	var passwordBody = map[string]string{
		"password": password,
	}

	reqBody, err := json.Marshal(passwordBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(userSetPasswordUrl, c.baseUrl, userId), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}

func (c *AcronisClient) ActivateWithMail(userId string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(userActivateUrl, c.baseUrl, userId), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}
	return nil
}

func (c *AcronisClient) FetchUser(userId string) (*apimodels.User, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(fetchUser, c.baseUrl, userId), nil)
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

	var user apimodels.User
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type offeringItemsResponse struct {
	Items []apimodels.OfferingItem `json:"items,omitempty"`
}

func (c *AcronisClient) GetOfferingItems(tenantId string) ([]apimodels.OfferingItem, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(enableOfferingItemsUrl, c.baseUrl, tenantId), nil)
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

	var respBody offeringItemsResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, err
	}
	return respBody.Items, nil
}

func (c *AcronisClient) SearchUserByLogin(login string) string {
	client, err := c.GetClient()
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(searchUrl, c.baseUrl, client.TenantId, login), nil)
	if err != nil {
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
