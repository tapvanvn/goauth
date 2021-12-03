package eth

import (
	"crypto/ecdsa"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/goauth/common"
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

	return crypto.Sign(message, client.privateKey)
}
func (client *OwnerClient) VerifySignature(message []byte, signature []byte) (bool, error) {

	return false, common.ErrNotImplement
}
