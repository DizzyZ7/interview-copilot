package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"interview-copilot/cli/internal/config"
)

type Client struct {
	BaseURL string
	Token   string
}

func New() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return &Client{BaseURL: cfg.BaseURL, Token: cfg.Token}, nil
}

func (c *Client) do(method, path string, body any, out any) error {
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}

	req, _ := http.NewRequest(method, c.BaseURL+path, &buf)
	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return errors.New(res.Status)
	}

	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}
	return nil
}

func (c *Client) SetToken(token string) error {
	cfg, _ := config.Load()
	cfg.Token = token
	return config.Save(cfg)
}
