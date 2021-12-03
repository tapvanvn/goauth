package eth

import "crypto/ecdsa"

type VerifyClient struct {
	publicKey *ecdsa.PublicKey
}
