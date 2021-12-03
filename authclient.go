package goauth

type IAuthClient interface {
	GetClientType() ClientType
	BeginSession(clientID AccountID, adapter IAdapter) (ISession, error) //frontend request to begin a signin process.
	GetBareResponse() IResponse                                          //response interface needed to verify
	GetBareSession() ISession
	Verify(session ISession, adapter IAdapter) (bool, error) //Verify if
}
