package eth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/goauth"
	"github.com/tapvanvn/goutil"
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

func (client *OwnerClient) GetEtherJSSignature(message []byte) ([]byte, error) {

	return goutil.EthersJSSignMessage(message, client.privateKey)
}

func (client *OwnerClient) GetSignature(message []byte) ([]byte, error) {

	signature, err := crypto.Sign(crypto.Keccak256(message), client.privateKey)
	if err != nil {
		return nil, err
	}
	signature[64] += 27 //So weir
	return signature, nil
}

func (client *OwnerClient) VerifyMessageSignature(message []byte, signature []byte) (bool, error) {

	hash := crypto.Keccak256(message)

	if signature[64] != 27 && signature[64] != 28 {

		fmt.Println("version:", signature[64])

		return false, goauth.ErrInvalidSignature
	}
	signature[64] -= 27

	pubKey, err := crypto.SigToPub(hash, signature)

	if err != nil {

		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	fmt.Println("address:", client.address, recoveredAddr)

	return client.address == recoveredAddr, nil
}

func (client *OwnerClient) VerifySignature(address string, title string, verifyMessage string, verifySignature []byte) (bool, error) {

	fromAddr := common.HexToAddress(address)

	doc := NewTypedDocument()
	doc.Parameters = append(doc.Parameters, &TypedParameter{
		Type:  "string",
		Name:  title,
		Value: verifyMessage,
	})

	if verifySignature[64] != 27 && verifySignature[64] != 28 {

		return false, goauth.ErrInvalidSignature
	}
	verifySignature[64] -= 27

	hash := doc.GetHash()

	pubKey, err := crypto.SigToPub(hash, verifySignature)

	if err != nil {

		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	return fromAddr.Hex() == recoveredAddr, nil
}

//MARK: implement IAuthClient
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

func (client *OwnerClient) VerifyAuthentication(clientID goauth.AccountID, response goauth.IResponse) (bool, error) {

	ethResponse := response.(*Response)

	verifySignature, err := hexutil.Decode(ethResponse.Signature)

	if err != nil {

		return false, err
	}
	if len(verifySignature) < 64 {

		return false, goauth.ErrInvalidSignature
	}

	return client.VerifySignature(string(clientID), ethResponse.MessageTitle, ethResponse.VerifyMessage, verifySignature)
}

func (client *OwnerClient) Verify(session goauth.ISession, response goauth.IResponse, adapter goauth.IAdapter) (bool, error) {

	ethSession := session.(*EthSession)
	ethResponse := response.(*Response)

	if ethSession == nil || ethResponse == nil {

		return false, goauth.ErrInvalidInformation
	}

	message := []byte(ethSession.SessionID)

	signature, err := hexutil.Decode(ethSession.VerifyMessage)

	if err != nil {

		return false, err
	}

	success, err := client.VerifyMessageSignature(message, signature)

	if !success || err != nil {

		return false, err
	}

	verifySignature, err := hexutil.Decode(ethResponse.Signature)

	if len(verifySignature) < 64 {
		return false, goauth.ErrInvalidSignature
	}
	return client.VerifySignature(ethSession.Address, ethResponse.MessageTitle, ethSession.VerifyMessage, verifySignature)
}

func (client *OwnerClient) ParseResponse(meta map[string]interface{}) (goauth.IResponse, error) {

	res := &Response{}
	infTitle, ok := meta["MessageTitle"]
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}
	title, ok := infTitle.(string)
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}
	res.MessageTitle = title

	infSignature, ok := meta["Signature"]
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}
	signature, ok := infSignature.(string)
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}
	res.Signature = signature

	if infVerifyMessage, ok := meta["VerifyMessage"]; ok {

		verifyMessage, ok := infVerifyMessage.(string)
		if !ok {
			return nil, goauth.ErrInvalidInformation
		}
		res.VerifyMessage = verifyMessage
	}

	return res, nil
}

func (client *OwnerClient) ParseSession(meta map[string]interface{}) (goauth.ISession, error) {

	session := &EthSession{}

	infSessionID, ok := meta["SessionID"]
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}

	sessionID, ok := infSessionID.(string)
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}

	session.SessionID = goauth.SessionID(sessionID)

	infAddress, ok := meta["Address"]
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}

	address, ok := infAddress.(string)
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}

	session.Address = address

	infVerifyMessage, ok := meta["VerifyMessage"]
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}

	verifyMessage, ok := infVerifyMessage.(string)
	if !ok {
		return nil, goauth.ErrInvalidInformation
	}

	session.VerifyMessage = verifyMessage

	return session, nil
}

//renew the session
func (client *OwnerClient) RenewSession(refreshToken string) (goauth.ISession, error) {
	return nil, goauth.ErrNotImplement
}
