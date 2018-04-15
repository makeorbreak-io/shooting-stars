package core

// Error string constants
var (
	// MessageOK is sent when a connection to web server is established
	MessageOK = "OK"
	// MessageDuel is sent when a duel started
	MessageDuel = "DUEL"
	// MessageClose is received when the client closes
	MessageClose = "CLOSE"
	// MessageClose is received when the client sends a ping
	MessagePing = "PING"
	// MessagePong is sent when the client sends a ping
	MessagePong = "PONG"
)
