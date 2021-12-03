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

	fmt.Println("--1")
	if client, ok := auth.clients[clientType]; ok {
		fmt.Println("--2")
		return client.BeginSession(clientAccountID, auth.adapter)
	}
	fmt.Println("--3")
	return nil, ErrClientNotFound
}

//when frontend process and send infomation for verifying account.
func (auth *Auth) VerifySession(session ISession, response IResponse) (bool, error) {

	clientType := session.GetClientType()

	if client, ok := auth.clients[clientType]; ok {

		return client.Verify(session, auth.adapter)
	}
	return false, nil
}
