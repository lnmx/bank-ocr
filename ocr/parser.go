package ocr

import (
	"fmt"
)

var ErrIllegible = fmt.Errorf("one or more illegible digits")

type Parser struct {
	// number of digits
	account_length int

	// width/height of a digit
	digit_width  int
	digit_height int

	// digit shapes to match against
	digits []string
}

func NewParser(account_length int) Parser {
	p := Parser{
		digit_width:    3,
		digit_height:   4,
		account_length: account_length,
	}

	digits := []string{
		" _     _  _     _  _  _  _  _ ",
		"| |  | _| _||_||_ |_   ||_||_|",
		"|_|  ||_  _|  | _||_|  ||_| _|",
		"                              ",
	}

	for i := 0; i < 10; i++ {
		p.digits = append(p.digits, p.extractChar(digits, i))
	}

	return p
}

func (p Parser) Parse(entry Entry) (acct Account, err error) {
	err = p.checkInput(entry)

	if err != nil {
		return
	}

	output := make([]byte, p.account_length)
	for i := 0; i < p.account_length; i++ {
		if ok, d := p.parseChar(entry, i); ok {
			output[i] = d
		} else {
			output[i] = '?'
			err = ErrIllegible
		}
	}

	return Account(output), err
}

func (p Parser) parseChar(entry Entry, n int) (ok bool, d byte) {
	char := p.extractChar(entry.Lines, n)

	for i, digit := range p.digits {
		if char == digit {
			return true, byte('0' + i)
		}
	}

	return false, '?'
}

func (p Parser) extractChar(lines []string, n int) (char string) {
	for i := 0; i < p.digit_height; i++ {
		start := n * p.digit_width
		end := (n + 1) * p.digit_width
		char += lines[i][start:end]
	}

	return char
}

func (p Parser) checkInput(entry Entry) (err error) {
	if len(entry.Lines) < p.digit_height {
		return fmt.Errorf("expected", p.digit_height, "lines in entry, got", len(entry.Lines))
	}

	line_length := p.digit_width * p.account_length
	for i := 0; i < p.digit_height; i++ {
		if len(entry.Lines[i]) < line_length {
			return fmt.Errorf("line", i, "wrong length, expected", line_length, "got", len(entry.Lines[i]))
		}
	}

	return nil
}
