package jpush

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jmind-systems/jpush/option"
)

const (
	maxIdleConn = 50
	timeOut     = 30 * time.Second

	scheme = "https"
	host   = "api.jpush.cn"

	pathPush = "/v3/push"
)

// Client is responsible for communication with jPush API.
type Client struct {
	appKey    string
	appSecret string

	scheme string
	host   string

	client *http.Client
}

// NewClient returns newly allocated jpush client.
func NewClient(appKey, appSecret string, opts ...option.ClientOption) (*Client, error) {
	client := Client{
		scheme:    scheme,
		host:      host,
		appKey:    appKey,
		appSecret: appSecret,
	}

	var settings option.DialSettings
	for _, opt := range opts {
		opt.Apply(&settings)
	}

	// Prepare http.Client.
	if settings.HTTPClient != nil {
		client.client = settings.HTTPClient
	} else {
		client.client = &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        maxIdleConn,
				MaxIdleConnsPerHost: maxIdleConn,
			},
			Timeout: timeOut,
		}
	}

	if settings.Endpoint != nil && strings.Contains(*settings.Endpoint, "://") {
		uri, err := url.Parse(*settings.Endpoint)
		if err != nil {
			return nil, fmt.Errorf("endpoint: %w", err)
		}

		client.scheme, client.host = uri.Scheme, uri.Host
	} else if settings.Endpoint != nil {
		client.scheme, client.host = scheme, *settings.Endpoint
	}

	return &client, nil
}

// Push dispatches new push request to jPush API.
func (c *Client) Push(ctx context.Context, payload *Request) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	uri := c.scheme + "://" + c.host + pathPush

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.appKey, c.appSecret)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("jpush: failed to send")
	}

	return nil
}
