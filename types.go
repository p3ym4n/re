package re

import (
	"encoding/json"
	"fmt"
)

// Kind indicates the kind of the error, default is KindUnexpected
type Kind string

func (k Kind) String() string {
	return string(k)
}

// Op is the identifier of the method
type Op string

func (op Op) String() string {
	return string(op)
}

// CodeInfo is added to the error for better observability
type CodeInfo struct {
	FileName     string `json:"file_name"`
	FunctionName string `json:"function_name"`
	LineNumber   int    `json:"line_number"`
}

func (ci *CodeInfo) String() string {
	return fmt.Sprintf("%s => %s line %d", ci.FileName, ci.FunctionName, ci.LineNumber)
}

// Meta is used for adding extra data to error
type Meta map[string]interface{}

func (m *Meta) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "metaData is not json serializable"
	}
	return string(bytes)
}

func (m *Meta) IsEmpty() bool {
	return len(*m) == 0
}
