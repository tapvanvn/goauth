package eth

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/goauth"
)

type StackVerifyClient struct {
	Addresses  map[string]string //map[user_name]address
	addressMux sync.Mutex
}

func (client *StackVerifyClient) AddAddress(username string, address string) {

	client.addressMux.Lock()
	client.Addresses[username] = address
	client.addressMux.Unlock()
}

func (client *StackVerifyClient) VerifySignature(doc goauth.Document, signature []byte) (bool, error) {

	//The signature used should be format like username.signature

	parts := strings.Split(string(signature), ".")
	if len(parts) != 2 {
		return false, goauth.ErrInvalidSignature
	}
	username := parts[0]

	client.addressMux.Lock()
	userAddress, hasUserAddress := client.Addresses[username]
	client.addressMux.Unlock()

	if !hasUserAddress {
		return false, goauth.ErrAccountNotFound
	}
	signatureBytes, err := hexutil.Decode(parts[1])

	if err != nil {
		return false, goauth.ErrInvalidSignature
	}

	if signatureBytes[64] != 27 && signatureBytes[64] != 28 {

		return false, goauth.ErrInvalidSignature
	}
	signatureBytes[64] -= 27

	hash := doc.GetHash()

	pubKey, err := crypto.SigToPub(hash, signatureBytes)

	if err != nil {

		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	return userAddress == recoveredAddr, nil
}

func (client *StackVerifyClient) GetClientType() goauth.ClientType {

	return goauth.ClientTypeEthereumStackVerify
}
