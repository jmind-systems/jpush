package jpush

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/jmind-systems/jpush/option"
)

const (
	maxIdleConn = 50
	timeOut     = 30 * time.Second
)

const (
	scheme = "https"
	host   = "api.jpush.cn"
)

const (
	pathPush = "/v3/push"
)

type Client struct {
	client *http.Client

	scheme string
	host   string

	appKey    string
	appSecret string
}

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

// TODO: Add retry feature.
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

	res, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(res))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.appKey, c.appSecret)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(res))

	if resp.StatusCode != http.StatusOK {
		return errors.New("jpush: failed to send")
	}

	return nil
}
