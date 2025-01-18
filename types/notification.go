package types

const (
	NotificationTypeError   NotificationTypeKey = "error"
	NotificationTypeNormal  NotificationTypeKey = "normal"
	NotificationTypeSuccess NotificationTypeKey = "success"
	NotificationTypeWarning NotificationTypeKey = "warning"
)

type NotificationTypeKey string

type Notification struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
