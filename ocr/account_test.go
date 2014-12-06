package ocr

import (
	"testing"
)

func TestAccountChecksum(t *testing.T) {

	length := 9

	valid := []Account{
		"711111111",
		"123456789",
		"490867715",
	}

	for _, acct := range valid {
		if !ValidateAccount(acct, length) {
			t.Fatal("should be valid:", acct)
		}
	}

	invalid := []Account{
		"888888888",
		"490067715",
		"012345678",
		"1",
		"012345?78",
	}

	for _, acct := range invalid {
		if ValidateAccount(acct, length) {
			t.Fatal("should be invalid:", acct)
		}
	}
}
