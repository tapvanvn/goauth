package common

type ClientType string

const (
	ClientTypeGoogle   = ClientType("google")
	ClientTypeApple    = ClientType("apple")
	ClientTypeEthereum = ClientType("eth")
	ClientTypeUserpass = ClientType("userpass")
	ClientTypeEmail    = ClientType("email")
	ClientTypePhone    = ClientType("phone")
)

type AccountID string
type SessionID string
type SessionAdapt interface{}
