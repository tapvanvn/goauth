package eth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
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
	fmt.Println("here 1")
	sessionID := adapter.NewSessionID()
	fmt.Println("here 2", sessionID)
	session := NewSession(sessionID, string(clientID))
	verifyMessage, err := client.GetSignature([]byte(sessionID))
	fmt.Println("here 3")
	if err != nil {
		return nil, err
	}

	session.verifyMessage = hexutil.Encode(verifyMessage)
	fmt.Println("here 4", session.verifyMessage)
	return session, nil
}
func (client *OwnerClient) Verify(session goauth.ISession, adapter goauth.IAdapter) (bool, error) {

	ethSession := session.(*EthSession)
	_ = ethSession
	return false, nil
}

func (client *OwnerClient) GetResponse() goauth.IResponse {
	return &Response{}
}
