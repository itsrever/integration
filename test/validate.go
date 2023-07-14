package test

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/stretchr/testify/require"
)

type JsonValidator struct {
	sch *jsonschema.Schema
}

func NewJsonValidator(jsonSchemaPath string) (*JsonValidator, error) {
	sch, err := jsonschema.Compile(jsonSchemaPath)
	if err != nil {
		return nil, err
	}
	return &JsonValidator{
		sch: sch,
	}, nil
}

func (v *JsonValidator) Validate(model string, data []byte) error {
	var mapData interface{}
	data = wrapWithModel(model, data)
	if err := json.Unmarshal(data, &mapData); err != nil {
		return err
	}
	return v.sch.Validate(mapData)
}

func (v *JsonValidator) RequireModel(t *testing.T, model string, reader io.ReadCloser) {
	data, err := io.ReadAll(reader)
	require.NoError(t, err)
	err = v.Validate(model, data)
	require.NoError(t, err)
}

func wrapWithModel(model string, data []byte) []byte {
	return []byte(fmt.Sprintf(`{"schemaVersion":"2020-12","%v":%v}`, model, string(data)))
}
