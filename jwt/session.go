package jwt

import "github.com/tapvanvn/goauth"

func NewSession(jwt goauth.SessionID, identifier string) *JWTSession {

	return &JWTSession{

		JWT:        jwt,
		Identifier: identifier,
	}
}

type JWTSession struct {
	JWT        goauth.SessionID `json:"SessionID"`
	Identifier string           `json:"Identifier"`
}

func (session *JWTSession) GetSessionID() goauth.SessionID {

	return session.JWT
}

func (session *JWTSession) GetClientAccountID() goauth.AccountID {

	return goauth.AccountID(session.Identifier)
}

func (session *JWTSession) GetClientType() goauth.ClientType {

	return goauth.ClientTypeJWT
}
