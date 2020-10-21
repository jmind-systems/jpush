package jpush

import "encoding/json"

type Platform string

const (
	Android Platform = "android"
	IOS     Platform = "ios"
)

type Request struct {
	// Required.
	Platform     Platform    `json:"platform"`               // Settings of push platform.
	Audience     Audience    `json:"audience"`               // Designation of push device.
	Message      *Message    `json:"message,omitempty"`      // Message content body.
	Notification interface{} `json:"notification,omitempty"` // Notification content body.
}

type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	Abtest         []string `json:"abtest,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
}

type Message struct {
	MsgContent  string          `json:"msg_content,omitempty"`
	Title       string          `json:"title,omitempty"`
	ContentType string          `json:"content_type,omitempty"`
	Extras      json.RawMessage `json:"extras,omitempty"`
}
