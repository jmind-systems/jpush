package payload

type AndroidNotification struct {
	Alert      string            `json:"alert"`
	Title      string            `json:"title,omitempty"`
	BuilderID  int               `json:"builder_id,omitempty"`
	ChannelID  string            `json:"channel_id,omitempty"`
	Priority   int               `json:"priority,omitempty"`
	Category   string            `json:"category,omitempty"`
	Style      int               `json:"style,omitempty"`
	AlertType  int               `json:"alert_type,omitempty"`
	BigText    string            `json:"big_text,omitempty"`
	Inbox      map[string]string `json:"inbox,omitempty"`
	BigPicPath string            `json:"big_pic_path,omitempty"`
	Extras     map[string]string `json:"extras,omitempty"`
	LargeIcon  string            `json:"large_icon,omitempty"`
	Intent     map[string]string `json:"intent,omitempty"`
}
