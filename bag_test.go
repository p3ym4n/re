package re

import (
	"errors"
	"reflect"
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
	if !reflect.DeepEqual(gotMeta,wantMeta){
		t.Errorf("the metas are not the same")
	}
}

func TestBag_AsMap(t *testing.T) {
	bag := New("first", errors.New("first err"), Meta{"this": "that"})
	mapped := bag.AsMap()

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
	if bag.codeInfo.FileName == "" || bag.codeInfo.FunctionName == "" || bag.codeInfo.LineNumber == 0 {
		t.Errorf("wrong codeInfo value for the bag")
	}
}