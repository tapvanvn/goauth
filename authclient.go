package goauth

type IAuthClient interface {
	GetClientType() ClientType
	BeginSession(clientID AccountID, adapter IAdapter) (ISession, error) //frontend request to begin a signin process.
	GetResponse() IResponse                                              //response interface needed to verify
	Verify(session ISession, adapter IAdapter) (bool, error)             //Verify if
}
