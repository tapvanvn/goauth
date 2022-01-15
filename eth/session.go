package eth

import "github.com/tapvanvn/goauth"

func NewSession(sessionID goauth.SessionID, address string) *EthSession {

	return &EthSession{

		SessionID: sessionID,
		Address:   address,
	}
}

type EthSession struct {
	SessionID     goauth.SessionID `json:"SessionID"`
	Address       string           `json:"Address"`
	VerifyMessage string           `json:"VerifyMessage"`
}

func (session *EthSession) GetSessionID() goauth.SessionID {
	return session.SessionID
}
func (session *EthSession) GetClientAccountID() goauth.AccountID {
	return goauth.AccountID(session.Address)
}
func (session *EthSession) GetClientType() goauth.ClientType {
	return goauth.ClientTypeEthereum
}
