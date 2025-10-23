package documents

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
)

func TestDocTypeImplementsInterfaces(t *testing.T) {
	var _ fmt.Stringer = (DocType)(-1)
	var _ json.Marshaler = (DocType)(-1)
	var _ json.Unmarshaler = (*DocType)(nil)
	var _ yaml.InterfaceMarshaler = (DocType)(-1)
	var _ yaml.InterfaceUnmarshaler = (*DocType)(nil)
	var _ driver.Valuer = (DocType)(-1)
	var _ sql.Scanner = (*DocType)(nil)
}

func TestMarshalJSON(t *testing.T) {
	testStruct := struct {
		DocType DocType `json:"doc_type" yaml:"doc_type"`
	}{
		DocType: DocBulletBackgroundPaper,
	}

	expectedJSON := `{"doc_type":"bullet-background-paper"}`
	actualJSON, err := json.Marshal(testStruct)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(actualJSON))
}

func TestMarshalYAML(t *testing.T) {
	testStruct := struct {
		DocType DocType `json:"doc_type" yaml:"doc_type"`
	}{
		DocType: DocBulletBackgroundPaper,
	}

	expectedYAML := "doc_type: bullet-background-paper\n"
	actualYAML, err := yaml.Marshal(testStruct)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(actualYAML))
}

func TestUnmarshalJSON(t *testing.T) {
	var testStruct = struct {
		DocType DocType `json:"doc_type" yaml:"doc_type"`
	}{}

	testJSON := `{
		"doc_type": "bullet-background-paper"
	}`

	err := json.Unmarshal([]byte(testJSON), &testStruct)
	assert.NoError(t, err)
	assert.Equal(t, DocBulletBackgroundPaper, testStruct.DocType)
}

func TestUnmarshalYAML(t *testing.T) {
	var testStruct = struct {
		DocType DocType `json:"doc_type" yaml:"doc_type"`
	}{}

	testYAML := `
doc_type: bullet-background-paper
`

	err := yaml.Unmarshal([]byte(testYAML), &testStruct)
	assert.NoError(t, err)
	assert.Equal(t, DocBulletBackgroundPaper, testStruct.DocType)
}
