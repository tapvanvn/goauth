package eth

import (
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type ClassParameter struct {
	Type  string
	Value interface{}
}

type ClassDocument struct {
	Parameters []*ClassParameter
}

func (doc *ClassDocument) GetHash() []byte {

	heads := []string{}
	values := []interface{}{}
	for _, param := range doc.Parameters {
		heads = append(heads, param.Type)
		values = append(values, param.Value)
	}

	return solsha3.SoliditySHA3(heads, values)
}
