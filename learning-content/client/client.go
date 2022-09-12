package learning_content_client

import (
	"errors"
	"fmt"
	"os"

	req "github.com/imroc/req/v3"
)

type Client struct {
	c *req.Client
}

type ClientOption func(c *Client)

func New(options ...ClientOption) *Client {
	var c = &Client{
		c: req.C(),
	}

	c.setApiURL(os.Getenv("LCMS_API_URL"))
	c.setApiKey(os.Getenv("LCMS_API_KEY"))

	for _, o := range options {
		o(c)
	}

	return c
}

func (c *Client) GetLatestRelease() (*Release, error) {
	var res, err = c.r().Get("/releases/latest")

	if err != nil {
		return nil, fmt.Errorf("error retrieving content: %w", err)
	}

	if !res.IsSuccess() {
		return nil, errors.New("error retrieving content: " + res.Status)
	}

	var r = new(Release)

	err = res.UnmarshalJson(r)
	if err != nil {
                return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
        }

	return r, nil
}

func (c *Client) r() *req.Request {
	return c.c.R()
}

func (c *Client) setApiURL(apiURL string) *Client {
	c.c.SetBaseURL(apiURL)
	return c
}

func WithApiURL(apiURL string) ClientOption {
	return func(c *Client) {
		c.setApiURL(apiURL)
	}
}

func (c *Client) setApiKey(apiKey string) *Client {
	c.c.SetCommonBearerAuthToken(apiKey)
	return c
}

func WithApiKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.setApiKey(apiKey)
	}
}
