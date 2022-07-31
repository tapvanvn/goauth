package momo

import "github.com/tapvanvn/goauth"

type MiniAppSession struct {
	SessionID     goauth.SessionID `json:"SessionID"`
	MiniAppUserID string           `json:"MiniAppUserID"`
}

func NewSession(sessionID goauth.SessionID, miniAppUserID string, authCode string) *MiniAppSession {
	return &MiniAppSession{
		SessionID:     sessionID,
		MiniAppUserID: miniAppUserID,
	}
}

func (session *MiniAppSession) GetSessionID() goauth.SessionID {
	return session.SessionID
}

func (session *MiniAppSession) GetClientAccountID() goauth.AccountID {
	return goauth.AccountID(session.MiniAppUserID)
}

func (session *MiniAppSession) GetClientType() goauth.ClientType {
	return goauth.ClientTypeMomoMiniapp
}
