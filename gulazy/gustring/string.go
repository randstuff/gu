package gustring

func Reverse(s string) string {
	var tReverse string
	for i := len(s) - 1; i >= 0; i-- {
		tReverse += string(s[i])
	}
	return tReverse
}

// from google github
// https://github.com/golang/example
//
func ReverseUsingRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
