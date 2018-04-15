package core

import "errors"

// Error string constants
var (
	// MessageOK is sent when a connection to web server is established
	MessageOK = "OK"
	// MessageDuel is sent when a duel started
	MessageDuel = errors.New("DUEL")
)
