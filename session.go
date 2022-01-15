package goauth

type ISession interface {
	GetSessionID() SessionID
	GetClientAccountID() AccountID
	GetClientType() ClientType
}

type IResponse interface {
}
