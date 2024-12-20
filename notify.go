package notify_go

type Notifier interface {
	Text(msg string) error
}
