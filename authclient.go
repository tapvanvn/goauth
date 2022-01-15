package goauth

type IAuthClient interface {
	GetClientType() ClientType

	BeginSession(clientID AccountID, adapter IAdapter) (ISession, error) //frontend request to begin a signin process.

	VerifyAuthentication(clientID AccountID, response IResponse) (bool, error) //verify authentication

	Verify(session ISession, response IResponse, adapter IAdapter) (bool, error) //Verify if

	//parse meta to response struct
	//ParseResponse(meta map[string]interface{}) (IResponse, error)

	//parse meta to session struct
	//ParseSession(meta map[string]interface{}) (ISession, error)

	//For the auth method that provice a renew machanism
	//If the auth method not support renew session, it should return not implement error
	RenewSession(refreshToken string) (ISession, error)
}
