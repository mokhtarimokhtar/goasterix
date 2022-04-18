package util

import (
	"encoding/hex"
	"strings"
)

/*
 Package util arrange implements some functions to manipulate strings for the testing.
*/

const (
	MsgFailInValue     = "FAIL: %s - Result %v - Expected %v"
	MsgSuccessInValue  = "SUCCESS: %s - Result %v - Expected %v"
	MsgFailInHex       = "FAIL: %s - Result %x - Expected %x"
	MsgSuccessInHex    = "SUCCESS: %s - Result %x - Expected %x"
	MsgFailInString    = "FAIL: %s - Result %x - Expected %x"
	MsgSuccessInString = "SUCCESS: %s - Result %x - Expected %x"
)

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
