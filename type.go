package goauth

type ClientType string
type SessionState int

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

type Document interface {
	GetHash() []byte
}

const (
	SessionStateInit = SessionState(0)
)
