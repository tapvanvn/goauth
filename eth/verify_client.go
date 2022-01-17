package eth

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/goauth"
)

func NewVerifyClient(walletAddress string) (*VerifyClient, error) {

	return &VerifyClient{
		address: walletAddress,
	}, nil
}

type VerifyClient struct {
	address string
	//publicKey *ecdsa.PublicKey
}

func (client *VerifyClient) VerifySignature(doc goauth.Document, signature []byte) (bool, error) {

	if signature[64] != 27 && signature[64] != 28 {

		return false, goauth.ErrInvalidSignature
	}
	signature[64] -= 27

	hash := doc.GetHash()

	pubKey, err := crypto.SigToPub(hash, signature)

	if err != nil {

		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	return client.address == recoveredAddr, nil
}

func (client *VerifyClient) GetClientType() goauth.ClientType {

	return goauth.ClientTypeEthereumVerify
}

func (client *VerifyClient) BeginSession(clientID goauth.AccountID, adapter goauth.IAdapter) (goauth.ISession, error) {

	sessionID := adapter.NewSessionID()

	return NewSession(sessionID, string(clientID)), nil
}

func (client *VerifyClient) Verify(session goauth.ISession, adapter goauth.IAdapter) (bool, error) {

	return false, nil
}

func (client *VerifyClient) RenewSession(refreshToken string) (goauth.ISession, error) {

	return nil, goauth.ErrNotImplement
}
