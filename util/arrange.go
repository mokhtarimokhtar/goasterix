package util

import (
	"encoding/hex"
	"strings"
)

/*
 Package util arrange implements some functions to manipulate strings for the testing.
*/

// HexStringToByte converts a hexadecimal string format to an array of byte.
// It is used to facilitate the testing.
func HexStringToByte(s string) ([]byte, error) {
	s = strings.ReplaceAll(s, " ", "")
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CleanStringMultiline cleans the escapes characters.
// It is used to facilitate the testing.
func CleanStringMultiline(s string) string {
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}
