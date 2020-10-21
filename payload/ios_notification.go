package payload

type IOSNotification struct {
	Alert            string            `json:"alert,omitempty"`
	Sound            string            `json:"sound,omitempty"`
	Badge            string            `json:"badge,omitempty"`
	ContentAvailable bool              `json:"content-available,omitempty"`
	MutableContent   bool              `json:"mutable-content,omitempty"`
	Category         string            `json:"category,omitempty"`
	Extras           map[string]string `json:"extras,omitempty"`
	ThreadId         string            `json:"thread-id,omitempty"`
}
