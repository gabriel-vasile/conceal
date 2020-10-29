package conceal

// ssn is used to scan a slice of bytes searching for SSNs.
// A SSN has 9 digits, formatted as 123-45-6789.
type ssn struct {
	// p1, p2, and p3 accumulate the digits of the SSN.
	p1 [3]byte
	p2 [2]byte
	p3 [4]byte
	// status tells how many digits of the SSN have been accumulated in p1, p2, and p3.
	status ssnStatus
}

type ssnStatus int

const (
	ssnStart ssnStatus = iota - 1
	ssnDigit1
	ssnDigit2
	ssnDigit3
	ssnDash1
	ssnDigit4
	ssnDigit5
	ssnDash2
	ssnDigit6
	ssnDigit7
	ssnDigit8
	ssnDigit9

	ssnLen = 11
)

// SSN finds all SSNs in s and returns a new slice with each occurence masked.
// If no number has been found, SSN returns s.
func SSN(s []byte) []byte {
	scanner := &ssn{status: ssnStart}
	return scanner.mask(s)
}

func (v *ssn) mask(s []byte) []byte {
	if len(s) < ssnLen {
		return s
	}
	masked := []byte{}
	for i := 0; i < len(s); i++ {
		v.status++
		c := s[i]
		// Look for digits and dashes.
		if !isDigit(c) && c != '-' {
			v.reset()
			continue
		}
		if v.status < ssnDash1 {
			if isDigit(c) {
				v.p1[v.status] = c
			} else {
				v.reset()
			}
		} else if v.status == ssnDash1 && c != '-' {
			v.reset()
		} else if v.status > ssnDash1 && v.status < ssnDash2 {
			if isDigit(c) {
				v.p2[v.status-ssnDash1-1] = c
			} else {
				v.reset()
			}
		} else if v.status == ssnDash2 && c != '-' {
			v.reset()
		} else if v.status > ssnDash2 && v.status <= ssnDigit9 {
			if isDigit(c) {
				v.p3[v.status-ssnDash2-1] = c
			} else {
				v.reset()
			}
		}
		if v.checkAcc() {
			if len(masked) == 0 {
				masked = append([]byte{}, s...)
			}
			maskSSN(masked, i)
			v.reset()
		}
	}

	if len(masked) > 0 {
		return masked
	}
	return s
}

// checkAcc returns true when v has accumulated the necessary number of digits
// for an SSN and the digits represent a valid SSN.
func (v *ssn) checkAcc() bool {
	if v.status != ssnDigit9 {
		return false
	}

	p1, p2, p3 := v.parts()

	if p1 == 0 || p1 == 666 || (p1 >= 900 && p1 <= 999) {
		return false
	}
	if p2 == 0 {
		return false
	}
	if p3 == 0 {
		return false
	}

	return true
}

func (v *ssn) parts() (int, int, int) {
	p1 := int(v.p1[0]-'0')*100 + int(v.p1[1]-'0')*10 + int(v.p1[2]-'0')
	p2 := int(v.p2[0]-'0')*10 + int(v.p2[1]-'0')
	p3 := int(v.p3[0]-'0')*1000 + int(v.p3[1]-'0')*100 + int(v.p3[2]-'0')*10 + int(v.p3[3]-'0')
	return p1, p2, p3
}

func (v *ssn) reset() {
	v.status = ssnStart
}

// maskSSN goes 11 bytes backwards and masks each byte.
func maskSSN(masked []byte, index int) {
	// reversed masked SSN number
	ssnMask := []byte("XXXX-XX-NSS")
	for i := index; i > index-ssnLen; i-- {
		masked[i] = ssnMask[index-i]
	}
}
