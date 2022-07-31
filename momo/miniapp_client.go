package momo

import (
	"github.com/tapvanvn/goauth"
)

type MiniappClient struct {
	AppID string
}

func NewMiniappClient(momoAppID string) (*MiniappClient, error) {
	return &MiniappClient{
		AppID: momoAppID,
	}, nil
}

//MARK: Implement IAuthClient
func (client *MiniappClient) GetClientType() goauth.ClientType {

	return goauth.ClientTypeMomoMiniapp
}

func (client *MiniappClient) BeginSession(clientID goauth.AccountID, adapter goauth.IAdapter) (goauth.ISession, error) { //frontend request to begin a signin process.

	sessionID := adapter.NewSessionID()

	session := NewSession(sessionID, string(clientID), "")

	return session, nil
}
func (client *MiniappClient) Authenticate(clientID goauth.AccountID, response goauth.IResponse) (bool, error) { //verify authentication
	//get authcode from response
	//check authcode from response
	return true, nil
}
func (client *MiniappClient) Verify(session goauth.ISession, response goauth.IResponse, adapter goauth.IAdapter) (bool, error) { //Verify session

	//
	return false, goauth.ErrNotImplement
}

//For the auth method that provice a renew machanism
//If the auth method not support renew session, it should return not implement error
func (client *MiniappClient) RenewSession(refreshToken string) (goauth.ISession, error) {

	return nil, goauth.ErrNotImplement
}
