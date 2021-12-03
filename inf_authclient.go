package goauth

import "github.com/tapvanvn/goauth/common"

type IAuthClient interface {
	GetClientType() common.ClientType
	GetAccountID(sessionAdapt common.SessionAdapt) common.AccountID                  //get account id in current provider system.
	BeginSession(id common.AccountID) (common.SessionID, common.SessionAdapt, error) //frontend request to begin a signin process.
	Verify(id common.AccountID, sessionAdapt common.SessionAdapt) error              //Verify if
}

type IRepo interface {
	NewSessionID() common.SessionID //issue new (alltime)unique sessionID
}
