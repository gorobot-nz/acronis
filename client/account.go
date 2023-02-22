package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorobot-nz/acronis/client/apimodels"
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

func (c *AcronisClient) CreateUser(accountCreate *apimodels.Account) (*apimodels.Account, error) {
	if isLogin := c.checkLogin(accountCreate.Login); !isLogin {
		return nil, errors.New("login is taken")
	}

	reqBody, err := json.Marshal(accountCreate)
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

	var account apimodels.Account
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *AcronisClient) ActivateWithPassword(accountId, password string) error {
	var passwordBody = map[string]string{
		"password": password,
	}

	reqBody, err := json.Marshal(passwordBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(userSetPasswordUrl, c.baseUrl, accountId), bytes.NewBuffer(reqBody))
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
		return errors.New("password was not set")
	}
	return nil
}

func (c *AcronisClient) ActivateWithMail(accountId string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(userActivateUrl, c.clientId, accountId), nil)
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
		return errors.New("mail was not send")
	}
	return nil
}
