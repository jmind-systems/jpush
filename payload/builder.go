// Package payload is a helper package which contains a payload
// builder to make constructing notification payloads easier.
package payload

type NotificationBuilder struct {
	notification *Notification
}

// NewNotification returns a new Payload struct
func NewNotification() *NotificationBuilder {
	return &NotificationBuilder{
		notification: &Notification{},
	}
}

type Notification struct {
	Alert   string               `json:"alert,omitempty"`
	Android *AndroidNotification `json:"android,omitempty"`
	IOS     *IOSNotification     `json:"ios,omitempty"`
}

func (n *NotificationBuilder) Android() *AndroidNotificationBuilder {
	if n.notification.Android == nil {
		n.notification.Android = &AndroidNotification{}
	}

	return &AndroidNotificationBuilder{notification: n.notification.Android}
}

func (n *NotificationBuilder) IOS() *IOSNotification {
	if n.notification.IOS != nil {
		return n.notification.IOS
	}

	return &IOSNotification{}
}

func (n *NotificationBuilder) Build() Notification {
	return *n.notification
}
