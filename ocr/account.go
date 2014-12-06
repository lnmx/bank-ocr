package ocr

// Account number (represented as a string of digit characters)
//
type Account string

func ValidateAccount(acct Account, length int) bool {
	if len(acct) != length {
		return false
	}

	total := 0
	for i, c := range acct {
		if c < '0' || c > '9' {
			return false
		}

		total += int(c-'0') * (len(acct) - i)
	}

	check := total % 11

	return check == 0
}
