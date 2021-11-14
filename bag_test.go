package re

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestBag_Chain(t *testing.T) {
	err1 := New("first", errors.New("first err"))
	err2 := err1.Chain("second")
	err3 := err2.Chain("third")
	want := []Op{"first", "second", "third"}
	got := err3.(*Bag).ops
	for i, op := range want {
		if got[i] != op {
			t.Error("the error operation is not correct")
		}
	}
}

func TestBag_ChainWithMeta(t *testing.T) {
	err1 := New("first", errors.New("first err"), Meta{"this": "that"})
	err2 := err1.Chain("second")
	err3 := err2.ChainWithMeta("third", Meta{"those": "them"})
	want := []Op{"first", "second", "third"}
	got := err3.(*Bag).ops
	for i, op := range want {
		if got[i] != op {
			t.Error("the error operation is not correct")
		}
	}

	wantMeta := Meta{
		"this": "that",
		"third": Meta{
			"those": "them",
		},
	}
	gotMeta := err3.(*Bag).metaData
	if !reflect.DeepEqual(gotMeta, wantMeta) {
		t.Errorf("the metas are not the same")
	}
}

func TestBag_RawMap(t *testing.T) {
	bag := New("first", errors.New("first err"), Meta{"this": "that"})
	mapped := bag.RawMap()

	if bag.kind.String() != mapped["kind"].(Kind).String() {
		t.Errorf("wrong kind for the bag")
	}
	if bag.message != mapped["message"] {
		t.Errorf("wrong message for the bag")
	}
	if bag.internal.Error() != mapped["internal"].(error).Error() {
		t.Errorf("wrong internal for the bag")
	}
	if !reflect.DeepEqual(bag.ops, mapped["operations"]) {
		t.Errorf("wrong ops for the bag")
	}
	gotMeta := mapped["meta_data"].(Meta)
	for k, v := range bag.metaData {
		if gotMeta[k] != v {
			t.Errorf("wrong metaData value for the bag")
		}
	}
	cf, ok := mapped["code_info"].(CodeInfo)
	if !ok {
		t.Errorf("wrong codeInfo value for the bag")
	}
	if cf.FileName == "" || cf.FunctionName == "" || cf.LineNumber == 0 {
		t.Errorf("wrong codeInfo value for the bag")
	}
}

func TestBag_ProcessedMap(t *testing.T) {

	t.Run("single depth", func(t *testing.T) {
		bag := New("first", errors.New("first err"), Meta{"this": "that"})
		processedMap := bag.ProcessedMap()

		if bag.kind.String() != processedMap["kind"] {
			t.Errorf("wrong kind for the bag")
		}
		if bag.internal.Error() != processedMap["internal"] {
			t.Errorf("wrong internal for the bag")
		}

		operations := make([]string, len(bag.ops))
		for i := len(bag.ops) - 1; i >= 0; i-- {
			operations = append(operations, bag.ops[i].String())
		}

		if processedMap["operations"] != strings.Join(operations, " => ") {
			t.Errorf("wrong ops for the bag")
		}
		marshalled, err := json.Marshal(bag.metaData)
		if err != nil {
			t.Fatal(err)
		}
		if string(marshalled) != processedMap["meta_data"] {
			t.Errorf("wrong meta_data for the bag")
		}
		if processedMap["code_info"] == "" {
			t.Errorf("wrong codeInfo value for the bag")
		}
	})

	t.Run("multi depth operations check", func(t *testing.T) {
		bag := New("first", errors.New("first err"), Meta{"this": "that"})
		bag.Chain("the second level")
		bag.Chain("the third")
		processedMap := bag.ProcessedMap()
		if processedMap["operations"] != "the third => the second level => first" {
			t.Errorf("wrong ops for the bag, got %s", processedMap["operations"])
		}
	})

}
