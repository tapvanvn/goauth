package goauth

type ClientType string

const (
	ClientTypeGoogle   = ClientType("google")
	ClientTypeApple    = ClientType("apple")
	ClientTypeEthereum = ClientType("eth")
	ClientTypeUserpass = ClientType("userpass")
	ClientTypeEmail    = ClientType("email")
	ClientTypePhone    = ClientType("phone")
	ClientTypeJWT      = ClientType("jwt")
)

type AccountID string
type SessionID string

type Document interface {
	GetHash() []byte
}
