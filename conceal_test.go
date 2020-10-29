package conceal_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/gabriel-vasile/conceal"
)

func TestConceal(t *testing.T) {
	txt := `
		Hi! 5555 5555 5555 4444 is my credit card number.
		My SSN is 289-17-5729 and my email is spam.me@email.com.
	`
	exp := `
		Hi! XXXX XXXX XXXX XXXX is my credit card number.
		My SSN is SSN-XX-XXXX and my email is sXXXXXe@eXXXl.com.
	`

	out := bytes.NewBuffer(nil)
	c := conceal.New(out, conceal.CardNumber, conceal.SSN, conceal.Email)

	if _, err := io.Copy(c, strings.NewReader(txt)); err != nil {
		t.Errorf("Conceal: %s", err)
	}

	if out.String() != exp {
		t.Logf("Exp: %s", exp)
		t.Logf("Got: %s", out.String())
		t.Fail()
	}
}
