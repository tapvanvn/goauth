package goauth

type ISession interface {
	GetSessionID() SessionID
	GetClientAccountID() AccountID
	GetClientType() ClientType
	GetState() SessionState
}

type IResponse interface {
}
