package jwt

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"

	"github.com/tapvanvn/goauth"
	"github.com/tapvanvn/godbengine/engine"
	"github.com/tapvanvn/goutil"
	"golang.org/x/oauth2/jws"
)

func NewClient(eng *engine.Engine, prefix string, tokenLifeTime time.Duration, refreshLifeTime time.Duration, ppk *rsa.PrivateKey) (*Client, error) {
	if eng.GetMemPool() == nil {
		return nil, goauth.ErrMempoolNotFound
	}
	//TODO: verify ppk
	client := &Client{
		dbEngine:        eng,
		tokenLifeTime:   tokenLifeTime,
		refreshLifeTime: refreshLifeTime,
		privateKey:      ppk,
	}
	if prefix != "" {
		client.prefix = fmt.Sprintf("%s_", prefix)
	}
	return client, nil
}

type Client struct {
	dbEngine        *engine.Engine
	prefix          string
	tokenLifeTime   time.Duration
	refreshLifeTime time.Duration
	privateKey      *rsa.PrivateKey
}

//MARK: implement IAuthClient
func (client *Client) GetClientType() goauth.ClientType {

	return goauth.ClientTypeJWT
}
func (client *Client) GetKey(refreshToken string) string {
	return fmt.Sprintf("%s%s", client.prefix, refreshToken)
}

//frontend request to begin a signin process.
func (client *Client) BeginSession(clientID goauth.AccountID, adapter goauth.IAdapter) (goauth.ISession, error) {

	memPool := client.dbEngine.GetMemPool()
	refreshToken := goutil.GenSecretKey(32)

	key := client.GetKey(refreshToken)
	//fmt.Println("key:", key, " clientID:", string(clientID))
	memPool.SetExpire(key, string(clientID), client.refreshLifeTime)

	aud := fmt.Sprintf("%s.%s", clientID, refreshToken)

	claim := &jws.ClaimSet{Iss: key, Aud: aud, Exp: time.Now().Unix() + int64(client.tokenLifeTime.Seconds())}
	header := &jws.Header{Algorithm: "HS256", Typ: "JWT"}
	jwt, err := jws.Encode(header, claim, client.privateKey)
	if err != nil {
		return nil, err
	}
	return NewSession(goauth.SessionID(jwt), string(clientID)), nil
}

//the verify authentication for jwt is not nessasary.
func (client *Client) VerifyAuthentication(clientID goauth.AccountID, response goauth.IResponse) (bool, error) {
	return false, goauth.ErrNotImplement
}

//verifying jwt
func (client *Client) Verify(session goauth.ISession, response goauth.IResponse, adapter goauth.IAdapter) (bool, error) {
	jwt := string(session.GetSessionID())
	if jwt == "" {

		return false, goauth.ErrInvalidInfomation
	}

	claim, err := jws.Decode(jwt)
	if err != nil {

		return false, goauth.ErrInvalidInfomation
	}
	if time.Now().Unix() > claim.Exp {

		return false, goauth.ErrSessionExpire
	}
	//k := claim.Iss
	parts := strings.Split(claim.Aud, ".")
	numParts := len(parts)
	if numParts < 2 {

		return false, goauth.ErrInvalidInfomation
	}
	refreshToken := parts[numParts-1]
	key := client.GetKey(refreshToken)
	memPool := client.dbEngine.GetMemPool()
	parts = parts[0 : numParts-1]
	jwtIdentifier, err := memPool.Get(key)

	if err != nil || jwtIdentifier != strings.Join(parts, ".") {

		return false, goauth.ErrInvalidInfomation
	}
	return true, nil
}

func (client *Client) RenewSession(refreshToken string) (goauth.ISession, error) {

	key := client.GetKey(refreshToken)

	memPool := client.dbEngine.GetMemPool()

	jwtIdentifier, err := memPool.Get(key)

	if err != nil || jwtIdentifier == "" {

		return nil, goauth.ErrInvalidInfomation
	}
	aud := fmt.Sprintf("%s.%s", jwtIdentifier, refreshToken)

	claim := &jws.ClaimSet{Iss: key, Aud: aud, Exp: time.Now().Unix() + int64(client.tokenLifeTime.Seconds())}
	header := &jws.Header{Algorithm: "HS256", Typ: "JWT"}
	jwt, err := jws.Encode(header, claim, client.privateKey)
	if err != nil {
		return nil, err
	}
	return NewSession(goauth.SessionID(jwt), string(jwtIdentifier)), nil

}
