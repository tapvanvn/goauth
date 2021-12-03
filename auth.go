package goauth

func NewAuth(adapter IAdapter) *Auth {
	return &Auth{
		adapter: adapter,
		clients: make(map[ClientType]IAuthClient),
	}
}

type Auth struct {
	clients map[ClientType]IAuthClient
	adapter IAdapter
}

func (auth *Auth) RegClient(client IAuthClient) {

	auth.clients[client.GetClientType()] = client
}

func (auth *Auth) UnregClient(clientType ClientType) {

	delete(auth.clients, clientType)
}

//start a session
func (auth *Auth) BeginSession(clientType ClientType, clientAccountID AccountID) (ISession, error) {

	if client, ok := auth.clients[clientType]; ok {

		return client.BeginSession(clientAccountID, auth.adapter)
	}

	return nil, ErrClientNotFound
}

//when frontend process and send infomation for verifying account.
func (auth *Auth) VerifySession(clientType ClientType, sessionMeta map[string]interface{}, responseMeta map[string]interface{}) (bool, error) {

	if client, ok := auth.clients[clientType]; ok {
		response, err := client.ParseResponse(responseMeta)
		if err != nil {
			return false, err
		}
		session, err := client.ParseSession(sessionMeta)
		if err != nil {
			return false, err
		}
		return client.Verify(session, response, auth.adapter)
	}
	return false, ErrClientNotFound
}
