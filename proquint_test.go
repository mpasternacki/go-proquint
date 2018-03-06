package proquint

import (
	"net"
	"testing"
)

var encodedLengths = []struct{ decoded, encoded int }{
	{0, 0},
	{1, 5},
	{2, 5},
	{3, 11},
	{4, 11},
	{6, 17},
}

func TestEncodedLen(t *testing.T) {
	for _, tc := range encodedLengths {
		if el := EncodedLen(tc.decoded); el != tc.encoded {
			t.Errorf("EncodedLen(%d) returned %d, expected %d",
				tc.decoded, el, tc.encoded)
		}
	}
}

func parseIPv4(s string) []byte {
	return net.ParseIP(s)[12:]
}

var encodedBytes = []struct {
	decoded []byte
	encoded string
}{
	// Examples provided in the paper
	{parseIPv4("127.0.0.1"), "lusab-babad"},
	{parseIPv4("63.84.220.193"), "gutih-tugad"},
	{parseIPv4("63.118.7.35"), "gutuk-bisog"},
	{parseIPv4("140.98.193.141"), "mudof-sakat"},
	{parseIPv4("64.255.6.200"), "haguz-biram"},
	{parseIPv4("128.30.52.45"), "mabiv-gibot"},
	{parseIPv4("147.67.119.2"), "natag-lisaf"},
	{parseIPv4("212.58.253.68"), "tibup-zujah"},
	{parseIPv4("216.35.68.215"), "tobog-higil"},
	{parseIPv4("216.68.232.21"), "todah-vobij"},
	{parseIPv4("198.81.129.136"), "sinid-makam"},
	{parseIPv4("12.110.110.204"), "budov-kuras"},
}

func TestEncodeToString(t *testing.T) {
	for _, tc := range encodedBytes {
		if enc := EncodeToString(tc.decoded); enc != tc.encoded {
			t.Errorf("EncodeToString(%v) returned %#v, expected %#v",
				tc.decoded, enc, tc.encoded)
		}
	}

}
