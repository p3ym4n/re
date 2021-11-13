package re

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	type test struct {
		op       Op
		internal error
		args     []interface{}
		result   map[string]interface{}
	}

	tests := []test{
		{
			op:       Op("first op"),
			internal: errors.New("the new err"),
			args:     []interface{}{},
			result: map[string]interface{}{
				"operations": []Op{"first op"},
				"internal":   errors.New("the new err"),
				"kind":       KindUnexpected.String(),
				"meta_data":  Meta{},
				"message":    "",
			},
		},
		{
			op:       Op("second op"),
			internal: errors.New("the new err2"),
			args:     []interface{}{KindInvalid , "the message to check"},
			result: map[string]interface{}{
				"operations": []Op{"second op"},
				"internal":   errors.New("the new err2"),
				"kind":       KindInvalid.String(),
				"meta_data":  Meta{"mamad" : "ali"},
				"message":    "the message to check",
			},
		},
	}

	for _, tc := range tests {
		bag := New(tc.op, tc.internal, tc.args...)
		if bag == nil {
			t.Fatal("expected the result not to be nil")
		}
		if bag.kind.String() != tc.result["kind"].(string) {
			t.Errorf("wrong kind for the bag")
		}
		if bag.message != tc.result["message"] {
			t.Errorf("wrong message for the bag")
		}
		if bag.internal.Error() != tc.result["internal"].(error).Error() {
			t.Errorf("wrong internal for the bag")
		}
		if !reflect.DeepEqual(bag.ops, tc.result["operations"]) {
			t.Errorf("wrong ops for the bag")
		}
		gotMeta := tc.result["meta_data"].(Meta)
		for k, v := range bag.metaData {
			if gotMeta[k] != v {
				t.Errorf("wrong metaData value for the bag")
			}
		}
		if bag.codeInfo.FileName == "" || bag.codeInfo.FunctionName == "" || bag.codeInfo.LineNumber == 0 {
			t.Errorf("wrong codeInfo value for the bag")
		}
	}
}
