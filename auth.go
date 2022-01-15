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

//when frontend process and send infomation for verifying the session.
func (auth *Auth) VerifySession(clientType ClientType, session ISession, response IResponse) (bool, error) {

	if client, ok := auth.clients[clientType]; ok {

		return client.Verify(session, response, auth.adapter)
	}

	return false, ErrClientNotFound
}

//when frontend process and send infomation for verifying the authentication of provider.
func (auth *Auth) VerifyAuthentication(clientType ClientType, clientAccountID AccountID, response IResponse) (bool, error) {

	if client, ok := auth.clients[clientType]; ok {

		return client.VerifyAuthentication(clientAccountID, response)
	}

	return false, ErrClientNotFound
}
