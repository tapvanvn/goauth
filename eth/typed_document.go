package eth

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type TypedParameter struct {
	Type  string
	Name  string
	Value string
}

type TypedDocument struct {
	Parameters []*TypedParameter
}

func (doc *TypedDocument) GetHash() []byte {

	heads := []string{}
	values := []string{}
	for _, param := range doc.Parameters {
		heads = append(heads, fmt.Sprintf("%s %s", param.Type, param.Name))
		values = append(values, param.Value)
	}
	head := strings.Join(heads, " ")
	value := strings.Join(values, " ")
	return crypto.Keccak256(crypto.Keccak256([]byte(head)), crypto.Keccak256([]byte(value)))

}
