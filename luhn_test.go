package conceal

import (
	"bytes"
	"testing"
)

func TestLuhn(t *testing.T) {
	tests := []struct {
		text string
		exp  string
	}{
		{
			"378282246310005",
			"XXXXXXXXXXXXXXX",
		},
		{
			"4222222222222",
			"XXXXXXXXXXXXX",
		},
		{
			"spaced cc 6011 1111 1111 1117",
			"spaced cc XXXX XXXX XXXX XXXX",
		},
		{
			"6011-1111-1111-1117dashed cc",
			"XXXX-XXXX-XXXX-XXXXdashed cc",
		},
		{
			"before378282246310005 after",
			"beforeXXXXXXXXXXXXXXX after",
		},
		{
			"invalid check digit 378282246310004",
			"invalid check digit 378282246310004",
		},

		// Various other CCs.
		{
			"3714496353984311",
			"XXXXXXXXXXXXXXX1",
		},
		{
			"378734493671000",
			"XXXXXXXXXXXXXXX",
		},
		{
			"5610591081018250",
			"XXXXXXXXXXXXXXXX",
		},
		{
			"30569309025904",
			"XXXXXXXXXXXXXX",
		},
		{
			"38520000023237",
			"XXXXXXXXXXXXXX",
		},
	}

	for _, test := range tests {
		got := CardNumber([]byte(test.text))
		if !bytes.Equal(got, []byte(test.exp)) {
			t.Errorf("Invalid CardNumber mask:\nExp: %s\nGot: %s\n", test.exp, got)
		}
	}
}
