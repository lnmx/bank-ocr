package main

import (
	"fmt"
	"github.com/lnmx/bank-ocr/ocr"
	"io"
	"os"
)

func main() {
	err := run()

	if err != nil {
		fmt.Println("error:", err)
	}
}

func run() error {

	entry_lines := 4
	acct_length := 9

	reader := ocr.NewEntryReader(entry_lines, os.Stdin)

	parser := ocr.NewParser(acct_length)

	for {

		// read Entry
		//
		entry, err := reader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("input read failed:", err)
		}

		// parse to Account
		//
		acct, err := parser.Parse(entry)

		if err == ocr.ErrIllegible {
			// ambiguous input
			//
			fmt.Println(acct, "ILL")
		} else if err != nil {
			// unexpected parse failure
			//
			return fmt.Errorf("parse failed: %s", err)
		} else if !ocr.ValidateAccount(acct, acct_length) {
			// checksum failure
			//
			output(acct, "ERR")
		} else {
			// OK
			//
			output(acct, "")
		}
	}

	return nil
}

func output(acct ocr.Account, status string) {
	fmt.Println(acct, status)
}
