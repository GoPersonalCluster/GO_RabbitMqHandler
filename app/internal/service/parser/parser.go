package parser

import (
	"encoding/json"
)

type Parser[T any] interface {
	Encode(data T) ([]byte, error)
	Decode(data []byte) (T, error)
}

type JsonParser[T any] struct {
}

func (p *JsonParser[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (p *JsonParser[T]) Decode(data []byte) (T, error) {
	var result T

	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
