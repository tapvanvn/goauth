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

func (client *OwnerClient) VerifyMessageSignature(message []byte, signature []byte) (bool, error) {

	hash := crypto.Keccak256(message)

	if signature[64] != 27 && signature[64] != 28 {

		return false, goauth.ErrInvalidSignature
	}
	signature[64] -= 27

	pubKey, err := crypto.SigToPub(hash, signature)

	if err != nil {

		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	return client.address == recoveredAddr, nil
}

func (client *OwnerClient) VerifySignature(address string, verifyMessage string, verifySignature []byte) (bool, error) {

	fromAddr := common.HexToAddress(address)

	doc := NewTypedDocument()
	doc.Parameters = append(doc.Parameters, &TypedParameter{
		Type:  "string",
		Name:  "Message",
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
func (client *OwnerClient) Verify(session goauth.ISession, response goauth.IResponse, adapter goauth.IAdapter) (bool, error) {

	ethSession := session.(*EthSession)
	ethResponse := response.(*Response)
	fmt.Println("step 0")
	if ethSession == nil || ethResponse == nil {

		return false, goauth.ErrInvalidInfomation
	}
	fmt.Println("step 1")
	message := crypto.Keccak256([]byte(ethSession.SessionID))

	signature, err := hexutil.Decode(ethSession.VerifyMessage)
	if err != nil {
		return false, err
	}
	fmt.Println("step 2")
	success, err := client.VerifyMessageSignature(message, signature)

	if !success || err != nil {

		return false, err
	}
	fmt.Println("step 3")
	verifySignature, err := hexutil.Decode(ethResponse.Signature)
	fmt.Println("step 3")
	return client.VerifySignature(ethSession.Address, ethSession.VerifyMessage, verifySignature)
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
		fmt.Println("fail on parse sessionID")
		return nil, goauth.ErrInvalidInfomation
	}

	sessionID, ok := infSessionID.(string)
	if !ok {
		fmt.Println("fail on parse sessionID 2")
		return nil, goauth.ErrInvalidInfomation
	}

	session.SessionID = goauth.SessionID(sessionID)

	infState, ok := meta["State"]
	if !ok {
		fmt.Println("fail on parse state")
		return nil, goauth.ErrInvalidInfomation
	}

	state := infState.(int)
	/*if !ok {
		fmt.Println("fail on parse state2")
		return nil, goauth.ErrInvalidInfomation
	}*/

	session.State = goauth.SessionState(state)

	infAddress, ok := meta["Address"]
	if !ok {
		fmt.Println("fail on parse address")
		return nil, goauth.ErrInvalidInfomation
	}

	address, ok := infAddress.(string)
	if !ok {
		fmt.Println("fail on parse address 2")
		return nil, goauth.ErrInvalidInfomation
	}

	session.Address = address

	infVerifyMessage, ok := meta["VerifyMessage"]
	if !ok {
		fmt.Println("fail on parse verify message")
		return nil, goauth.ErrInvalidInfomation
	}

	verifyMessage, ok := infVerifyMessage.(string)
	if !ok {
		fmt.Println("fail on parse verifymessage 2")
		return nil, goauth.ErrInvalidInfomation
	}

	session.VerifyMessage = verifyMessage

	return session, nil
}
