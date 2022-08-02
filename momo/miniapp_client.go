package momo

import (
	"fmt"
	"time"

	"github.com/tapvanvn/goauth"
	"github.com/tapvanvn/gomomo/config"
	"github.com/tapvanvn/gomomo/miniapp"
)

type MiniappClient struct {
	AppID      string
	momoClient *miniapp.MiniAppClient
}

func NewMiniappClient(momoAppID string, isDev bool, openSecret string, openPrivateKey string, openPublicKey string) (*MiniappClient, error) {
	momoConfig := &config.ClientConfig{
		OpenSecret:     openSecret,
		OpenPrivateKey: openPrivateKey,
		OpenPublicKey:  openPublicKey,
	}
	authClient := &MiniappClient{
		AppID:      momoAppID,
		momoClient: miniapp.NewMiniAppClient(isDev, momoConfig),
	}

	return authClient, nil
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
	appResponse := response.(*MiniAppAuthResponse)
	if appResponse == nil {

		return false, goauth.ErrInvalidInformation
	}
	accessToken, err := client.momoClient.RequestAccessToken(appResponse.PartnerUserID, appResponse.AuthCode)
	if err != nil {
		return false, err
	}
	fmt.Println("momo access Token", accessToken.Token, "ex:", accessToken.ExpiredTime)
	return accessToken.ExpiredTime > time.Now().Unix(), nil
}
func (client *MiniappClient) Verify(session goauth.ISession, response goauth.IResponse, adapter goauth.IAdapter) (bool, error) { //Verify session

	return false, goauth.ErrNotImplement
}

//For the auth method that provice a renew machanism
//If the auth method not support renew session, it should return not implement error
func (client *MiniappClient) RenewSession(refreshToken string) (goauth.ISession, error) {

	return nil, goauth.ErrNotImplement
}
