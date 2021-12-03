package goauth

import "github.com/tapvanvn/goauth/common"

func NewAuth(repo IRepo) *Auth {
	return &Auth{
		repo:    repo,
		clients: make(map[common.ClientType]IAuthClient),
	}
}

type Auth struct {
	clients map[common.ClientType]IAuthClient
	repo    IRepo
}

func (auth *Auth) RegClient(client IAuthClient) {

	auth.clients[client.GetClientType()] = client
}

func (auth *Auth) UnregClient(clientType common.ClientType) {

	delete(auth.clients, clientType)
}
