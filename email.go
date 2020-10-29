package conceal

import (
	"bytes"
	"regexp"
)

var emailReg = regexp.MustCompile("(?i)([A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,24})")

// Email masks all email addresses from the input.
func Email(in []byte) []byte {
	return emailReg.ReplaceAllFunc(in, maskEmail)
}

func maskEmail(email []byte) []byte {
	atByte := byte('@')
	dotByte := byte('.')
	at := bytes.LastIndex(email, []byte{atByte})
	username, domain := email[:at], email[at+1:]

	dot := bytes.LastIndex(domain, []byte{dotByte})
	dName, dom := domain[:dot], domain[dot+1:]

	ret := append([]byte{}, maskEmailComponent(username)...)
	ret = append(ret, atByte)
	ret = append(ret, maskEmailComponent(dName)...)
	ret = append(ret, dotByte)
	ret = append(ret, dom...)
	return ret
}

// maskEmailComponent replaces all but the first and the last character from in
// with 'X'. When the input is shorter than 5 characters the first and last
// character are replaced too.
func maskEmailComponent(in []byte) []byte {
	l := len(in)
	x := []byte{byte('X')}
	if l < 5 {
		return bytes.Repeat(x, l)
	}

	ret := append([]byte{in[0]}, bytes.Repeat(x, l-2)...)
	ret = append(ret, in[l-1])
	return ret
}
