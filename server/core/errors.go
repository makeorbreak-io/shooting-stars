package core

import "errors"

// Error string constants
var (
	// ErrorBadRequest is thrown when a user makes a bad request
	ErrorBadRequest = errors.New("BAD_REQUEST")
	// ErrorDuplicated is thrown when a user tries to create an
	// object that already exists
	ErrorDuplicated = errors.New("DUPLICATED")
	// ErrorInternalServerError is thrown when a an internal server error occurs
	ErrorInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	// ErrorNotFound is thrown when a user tries to access an object
	// that does not exist
	ErrorNotFound = errors.New("NOT_FOUND")
	// ErrorBadLogin is thrown when the login credentials are invalid
	ErrorBadLogin = errors.New("BAD_LOGIN")
)
