package eth

import "github.com/tapvanvn/goauth"

func NewSession(sessionID goauth.SessionID, address string) *EthSession {

	return &EthSession{

		sessionID: sessionID,
		state:     goauth.SessionStateInit,
		address:   address,
	}
}

type EthSession struct {
	sessionID     goauth.SessionID
	state         goauth.SessionState
	address       string
	verifyMessage string
}

func (session *EthSession) GetSessionID() goauth.SessionID {
	return session.sessionID
}
func (session *EthSession) GetClientAccountID() goauth.AccountID {
	return goauth.AccountID(session.address)
}
func (session *EthSession) GetClientType() goauth.ClientType {
	return goauth.ClientTypeEthereum
}
func (session *EthSession) GetState() goauth.SessionState {
	return session.state
}
