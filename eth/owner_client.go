package eth

import (
	"crypto/ecdsa"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/goauth"
)

func NewOwnerClient(privateKeyShadow string) (*OwnerClient, error) {

	if strings.HasPrefix(privateKeyShadow, "0x") {

		privateKeyShadow = privateKeyShadow[2:]
	}
	privateKey, err := crypto.HexToECDSA(privateKeyShadow)

	if err != nil {

		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {

		return nil, errors.New("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &OwnerClient{

		address:    address,
		privateKey: privateKey,
		publicKey:  publicKeyECDSA,
	}, nil
}

type OwnerClient struct {
	address    string
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func (client *OwnerClient) GetSignature(message []byte) ([]byte, error) {

	return crypto.Sign(crypto.Keccak256(message), client.privateKey)
}

func (client *OwnerClient) VerifySignature(message []byte, signature []byte) (bool, error) {

	return false, goauth.ErrNotImplement
}

func (client *OwnerClient) GetClientType() goauth.ClientType {

	return goauth.ClientTypeEthereum
}

func (client *OwnerClient) BeginSession(clientID goauth.AccountID, adapter goauth.IAdapter) (goauth.ISession, error) {

	sessionID := adapter.NewSessionID()

	session := NewSession(sessionID, string(clientID))
	verifyMessage, err := client.GetSignature([]byte(sessionID))

	if err != nil {
		return nil, err
	}

	session.VerifyMessage = hexutil.Encode(verifyMessage)

	return session, nil
}
func (client *OwnerClient) Verify(session goauth.ISession, adapter goauth.IAdapter) (bool, error) {

	ethSession := session.(*EthSession)
	_ = ethSession
	return false, nil
}

func (client *OwnerClient) ParseResponse(meta map[string]interface{}) (goauth.IResponse, error) {

	res := &Response{}
	infSignature, ok := meta["Signature"]
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}
	signature, ok := infSignature.(string)
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}
	res.Signature = signature
	return res, nil
}

func (client *OwnerClient) ParseSession(meta map[string]interface{}) (goauth.ISession, error) {

	session := &EthSession{}

	infSessionID, ok := meta["SessionID"]
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}

	sessionID, ok := infSessionID.(string)
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}

	session.SessionID = goauth.SessionID(sessionID)

	infState, ok := meta["State"]
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}

	state, ok := infState.(goauth.SessionState)
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}

	session.State = state

	infAddress, ok := meta["Address"]
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}

	address, ok := infAddress.(string)
	if !ok {
		return nil, goauth.ErrInvalidInfomation
	}

	session.Address = address

	return session, nil
}
