package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/goauth/common"
)

type Client struct {
	address string

	publicKey *ecdsa.PublicKey
}

func (client *Client) VerifySignature(doc common.Document, signature []byte) (bool, error) {

	if signature[64] != 27 && signature[64] != 28 {

		return false, common.ErrInvalidSignature
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
