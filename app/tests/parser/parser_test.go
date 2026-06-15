package parser

import (
	"go_rabbitmqhandler/internal/service/parser"
	"reflect"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonParserEncode(t *testing.T) {
	parser := parser.JsonParser[Person]{}

	person := Person{
		Name: "Walter",
		Age:  30,
	}

	result, err := parser.Encode(person)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `{"name":"Walter","age":30}`

	if string(result) != expected {
		t.Errorf("expected %s, got %s", expected, string(result))
	}
}

func TestJsonParserDecode(t *testing.T) {
	parser := parser.JsonParser[Person]{}

	data := []byte(`{"name":"Walter","age":30}`)

	result, err := parser.Decode(data)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := Person{
		Name: "Walter",
		Age:  30,
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %+v, got %+v", expected, result)
	}
}

func TestJsonParserDecodeInvalidJson(t *testing.T) {
	parser := parser.JsonParser[Person]{}

	data := []byte(`{"name":"Walter","age":}`)

	_, err := parser.Decode(data)

	if err == nil {
		t.Error("expected an error, got nil")
	}
}
