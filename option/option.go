package option

import (
	"net/http"
)

type Credentials struct {
	AppKey    string
	AppSecret string
}

type DialSettings struct {
	Endpoint    *string
	Credentials *Credentials
	HTTPClient  *http.Client
}

type ClientOption interface {
	Apply(*DialSettings)
}

type ClientOptionFunc func(*DialSettings)

func (f ClientOptionFunc) Apply(settings *DialSettings) { f(settings) }

// WithEndpoint returns a ClientOption that overrides the default endpoint to be used for a service.
func WithCredentials(appKey, appSecret string) ClientOption {
	return ClientOptionFunc(func(settings *DialSettings) {
		settings.Credentials.AppKey = appKey
		settings.Credentials.AppSecret = appSecret
	})
}

// WithEndpoint returns a ClientOption that overrides the default endpoint to be used for a service.
func WithEndpoint(url string) ClientOption {
	return ClientOptionFunc(func(settings *DialSettings) {
		settings.Endpoint = &url
	})
}

// WithHTTPClient returns a ClientOption that specifies the HTTP client to use
// as the basis of communications. This option may only be used with services
// that support HTTP as their communication transport. When used, the
// WithHTTPClient option takes precedent over all other supplied options.
func WithHTTPClient(client *http.Client) ClientOption {
	return ClientOptionFunc(func(settings *DialSettings) {
		settings.HTTPClient = client
	})
}
