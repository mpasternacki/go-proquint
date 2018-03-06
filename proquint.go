package proquint

import "strings"

var proquint_consonants = []rune("bdfghjklmnprstvz")
var proquint_vowels = []rune("aiou")

// DecodedLen returns the length of an encoding of n source bytes.
func EncodedLen(n int) (rv int) {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 5
	}
	return ((n+1)/2)*6 - 1
}

// EncodeToString returns proquint encoding of src as a string.
func EncodeToString(src []byte) string {
	if len(src)&1 != 0 {
		src = append([]byte{0}, src...)
	}
	res := make([]string, len(src)/2)
	for i := 0; i < len(src); i += 2 {
		res[i/2] = string([]rune{
			proquint_consonants[src[i]&0xf0>>4],
			proquint_vowels[src[i]&0x0c>>2],
			proquint_consonants[(src[i]&0x03<<2)|(src[i+1]&0xc0>>6)],
			proquint_vowels[src[i+1]&0x30>>4],
			proquint_consonants[src[i+1]&0x0f],
		})
	}
	return strings.Join(res, "-")
}
