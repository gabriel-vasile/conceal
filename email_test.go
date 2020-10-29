package conceal

import (
	"bytes"
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		text string
		exp  string
	}{
		{
			`hello.world@example.com`,
			`hXXXXXXXXXd@eXXXXXe.com`,
		},
		{
			`text before hello.world@example.com text after`,
			`text before hXXXXXXXXXd@eXXXXXe.com text after`,
		},
		{
			`hi@example.com`,
			`XX@eXXXXXe.com`,
		},
		{
			`text before hi@example.com text after`,
			`text before XX@eXXXXXe.com text after`,
		},
		{
			`hi@shrt.com`,
			`XX@XXXX.com`,
		},
		{
			`text before hi@shrt.com text after`,
			`text before XX@XXXX.com text after`,
		},
	}
	for _, test := range tests {
		if got := Email([]byte(test.text)); !bytes.Equal(got, []byte(test.exp)) {
			t.Errorf("Invalid Email mask:\nExp: %s\nGot: %s\n", test.exp, got)
		}
	}
}
