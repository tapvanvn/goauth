package goauth

type IAuthClient interface {
	GetClientType() ClientType
	Authenticate(clientID AccountID, response IResponse) (bool, error) //verify authentication

	BeginSession(clientID AccountID, adapter IAdapter) (ISession, error)         //frontend request to begin a signin process.
	Verify(session ISession, response IResponse, adapter IAdapter) (bool, error) //Verify session
	//For the auth method that provice a renew machanism
	//If the auth method not support renew session, it should return not implement error
	RenewSession(refreshToken string) (ISession, error)
}
