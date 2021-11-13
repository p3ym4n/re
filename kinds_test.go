package re

import (
	"errors"
	"net/http"
	"testing"
)

func TestAddKind(t *testing.T) {

	const KindTesting Kind = "testing"

	AddKind(KindTesting, 200)

	got, has := kinds[KindTesting]
	if !has {
		t.Errorf("excpexted to has the %s as a kind", KindTesting)
	}
	if got != 200 {
		t.Errorf("excpexted to has %v as the value of kind %s", 200, KindTesting)
	}

}

func TestHttpCode(t *testing.T) {

	const KindOutOfNowhere Kind = "out of no where"
	bag := New("testing operation", errors.New("no new err"), KindOutOfNowhere)
	gotCode := HttpCode(bag)
	if gotCode != http.StatusInternalServerError {
		t.Errorf("non declared kind wrong http code")
	}

	bag2 := New("testing operation", errors.New("no new err"))
	gotCode2 := HttpCode(bag2)
	if gotCode2 != http.StatusInternalServerError {
		t.Errorf("default kind wrong http code")
	}

	bag3 := New("testing operation", errors.New("no new err"), KindForbidden)
	gotCode3 := HttpCode(bag3)
	if gotCode3 != http.StatusForbidden {
		t.Errorf("declared kind wrong http code")
	}

}
