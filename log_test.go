package re

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {

	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	Log(New("testing operation", errors.New("no new err")))
	res := buf.String()
	if !strings.ContainsAny(res , "testing operation"){
		t.Errorf("the error message should be in the logs")
	}
	if !strings.ContainsAny(res , "no new err"){
		t.Errorf("the internal error should be in the logs")
	}
}
