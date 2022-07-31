package goauth

type ClientType string

const (
	ClientTypeUnknown             = ClientType("unknown")
	ClientTypeGoogle              = ClientType("google")
	ClientTypeApple               = ClientType("apple")
	ClientTypeEthereum            = ClientType("eth")
	ClientTypeEthereumVerify      = ClientType("eth_verify")
	ClientTypeEthereumStackVerify = ClientType("eth_stack_verify")
	ClientTypeUserpass            = ClientType("userpass")
	ClientTypeEmail               = ClientType("email")
	ClientTypePhone               = ClientType("phone")
	ClientTypeJWT                 = ClientType("jwt")
	ClientTypeMomoMiniapp         = ClientType("momo_miniapp")
)

type AccountID string
type SessionID string

type Document interface {
	GetHash() []byte
}
