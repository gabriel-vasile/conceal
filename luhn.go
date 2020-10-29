package conceal

const (
	luhnMinLen = 13
	luhnMaxLen = 16
	capacity   = luhnMaxLen << 1
)

var (
	doubledAndSummed = [10]byte{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
)

type luhn struct {
	end, length           int
	currentSums, nextSums [capacity]byte
}

func CardNumber(s []byte) []byte {
	scanner := &luhn{}
	return scanner.mask(s)
}

func (l *luhn) mask(s []byte) []byte {
	if len(s) < luhnMinLen {
		return s
	}

	l.reset()
	masked := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isDigit(c) {
			l.push(c)
			if matchingDigits := l.matchingDigits(); matchingDigits > 0 {
				if len(masked) == 0 {
					masked = append([]byte{}, s...)
				}
				mask(masked, i, matchingDigits)
			}
		} else if !isSeparator(c) {
			l.reset()
		}
	}
	if len(masked) == 0 {
		return s
	}
	return masked
}

func (l *luhn) push(digit byte) {
	l.swapArrays()
	val := digit - '0'
	l.accumulateCurr(val)
	l.accumulateNext(doubledAndSummed[val])
	l.end = wrap(l.end + 1)
	if l.length < luhnMaxLen {
		l.length++
	}
}

func (l *luhn) accumulateCurr(val byte) {
	l.currentSums[l.end] = l.currentSums[wrap(l.end-1)] + val
}
func (l *luhn) accumulateNext(val byte) {
	l.nextSums[l.end] = l.nextSums[wrap(l.end-1)] + val
}
func (l *luhn) swapArrays() {
	l.currentSums, l.nextSums = l.nextSums, l.currentSums
}

func (l *luhn) matchingDigits() int {
	for llen := l.length; llen >= luhnMinLen; llen-- {
		last := wrap(l.end - 1)
		base := l.currentSums[wrap(last-l.length)]
		sum := l.currentSums[last]
		if (sum-base)%10 == 0 {
			return llen
		}
	}
	return 0
}

func (l *luhn) reset() {
	l.length = 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSeparator(c byte) bool {
	return c == ' ' || c == '-'
}

func wrap(i int) int {
	return i & (capacity - 1)
}

func mask(masked []byte, last, count int) {
	for count > 0 {
		c := masked[last]
		if isDigit(c) {
			masked[last] = 'X'
			count--
		} else if c == 'X' {
			count--
		}
		last--
	}
}
