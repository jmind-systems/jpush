package jpush

import "encoding/json"

// Platform represents target client OS.
type Platform string

const (
	// Android is an Android operating system.
	Android Platform = "android"
	// IOS is an iOS operating system.
	IOS Platform = "ios"
)

// Request represents push notification request.
type Request struct {
	Platform     Platform    `json:"platform"`
	Audience     Audience    `json:"audience"`
	Message      *Message    `json:"message,omitempty"`
	Notification interface{} `json:"notification,omitempty"`
}

// Audience represents push notification message audience.
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	Abtest         []string `json:"abtest,omitempty"`
	RegistrationID []string `json:"registration_id,omitempty"`
}

// Message represents request push notification message entity.
type Message struct {
	MsgContent  string          `json:"msg_content,omitempty"`
	Title       string          `json:"title,omitempty"`
	ContentType string          `json:"content_type,omitempty"`
	Extras      json.RawMessage `json:"extras,omitempty"`
}
