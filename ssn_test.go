package conceal

import (
	"bytes"
	"testing"
)

func TestSSN(t *testing.T) {
	tests := []struct {
		text string
		exp  string
	}{
		{
			"123-45-6789",
			"SSN-XX-XXXX",
		},
	}

	for _, test := range tests {
		got := SSN([]byte(test.text))
		if !bytes.Equal(got, []byte(test.exp)) {
			t.Errorf("Invalid SSN mask:\nExp: %s\nGot: %s\n", test.exp, got)
		}
	}
}
