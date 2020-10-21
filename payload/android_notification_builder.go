package payload

type AndroidNotificationBuilder struct {
	notification *AndroidNotification
}

// NewNotification returns a new Payload struct
func NewAndroidNotification() *AndroidNotificationBuilder {
	return &AndroidNotificationBuilder{
		notification: &AndroidNotification{},
	}
}
