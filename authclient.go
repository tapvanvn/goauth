package goauth

type IAuthClient interface {
	GetClientType() ClientType
	BeginSession(clientID AccountID, adapter IAdapter) (ISession, error) //frontend request to begin a signin process.

	Verify(session ISession, adapter IAdapter) (bool, error) //Verify if
	ParseResponse(meta map[string]interface{}) (IResponse, error)
	ParseSession(meta map[string]interface{}) (ISession, error)
}
