package eth

import (
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func NewClassicDocument() *ClassicDocument {
	return &ClassicDocument{
		Parameters: make([]*ClassicParameter, 0),
	}
}

type ClassicParameter struct {
	Type  string
	Value interface{}
}

type ClassicDocument struct {
	Parameters []*ClassicParameter
}

func (doc *ClassicDocument) GetHash() []byte {

	heads := []string{}
	values := []interface{}{}
	for _, param := range doc.Parameters {
		heads = append(heads, param.Type)
		values = append(values, param.Value)
	}

	return solsha3.SoliditySHA3(heads, values)
}

func (doc *ClassicDocument) Append(params ...*ClassicParameter) {

	doc.Parameters = append(doc.Parameters, params...)
}
