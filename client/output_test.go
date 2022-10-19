package client

import (
	"testing"
)

func TestOutputFormatting(t *testing.T) {

	serializedJson := `{"txguid":"test234-2-234234-2342f2ddd2-234234","numbers":"4474123456789","smsparts":"2","encoding":"utf-8"}`
	serializedYaml := `txguid: test234-2-234234-2342f2ddd2-234234
numbers: "4474123456789"
smsparts: "2"
encoding: utf-8
`

	response := &SuccessResponse{
		TxGuid:   "test234-2-234234-2342f2ddd2-234234",
		Numbers:  "4474123456789",
		SmsParts: "2",
		Encoding: "utf-8",
	}

	result, err := Output(response, "json")

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result != serializedJson {
		t.Errorf("Should be serialized correctly but got '%v'", result)
	}

	result, err = Output(response, "yaml")

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result != serializedYaml {
		t.Errorf("Should be serialized correctly but got '%v'", result)
	}

}
