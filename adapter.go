package goauth

type IAdapter interface {
	NewSessionID() SessionID //create
}
