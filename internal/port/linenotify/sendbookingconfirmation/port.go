package sendbookingconfirmation

// Messenger is the port booking confirmation depends on to push a LINE message.
type Messenger interface {
	Push(to, message string) error
}
