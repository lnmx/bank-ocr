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

	reader := ocr.NewEntryReader(4, os.Stdin)

	parser := ocr.NewParser()

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

		if err != nil {
			return fmt.Errorf("parse failed: %s", err)
		}

		// output
		//
		fmt.Println(acct)
	}

	return nil
}
