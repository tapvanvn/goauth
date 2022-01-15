package goauth

type IAdapter interface {
	NewSessionID() SessionID //issue a new sessionID
}
