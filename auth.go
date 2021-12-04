package goauth

import "fmt"

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
	fmt.Println(" auth VerifySession 1")
	if client, ok := auth.clients[clientType]; ok {
		response, err := client.ParseResponse(responseMeta)
		fmt.Println(" auth VerifySession 2")
		if err != nil {
			return false, err
		}
		fmt.Println(" auth VerifySession 3")
		session, err := client.ParseSession(sessionMeta)
		if err != nil {
			fmt.Println(" auth VerifySession 3", err)
			return false, err
		}
		fmt.Println(" auth VerifySession 4")
		return client.Verify(session, response, auth.adapter)
	}
	fmt.Println(" auth VerifySession 5")
	return false, ErrClientNotFound
}
