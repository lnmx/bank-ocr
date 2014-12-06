package ocr

import (
	"bufio"
	"io"
)

// EntryReader reads entries consisting of linecount lines from input.
//
type EntryReader struct {
	linecount int
	input     io.Reader
	scanner   *bufio.Scanner
}

func NewEntryReader(linecount int, input io.Reader) *EntryReader {
	return &EntryReader{
		linecount: linecount,
		input:     input,
		scanner:   bufio.NewScanner(input),
	}
}

// Return the next set of lines, or an error.
//
func (r *EntryReader) Next() (entry Entry, err error) {
	lines := []string{}

	for i := 0; i < r.linecount; i++ {

		if !r.scanner.Scan() {
			if err = r.scanner.Err(); err != nil {
				// input error
				return
			} else if i == 0 {
				// EOF between entries -- OK
				err = io.EOF
				return
			} else {
				// EOF in the middle of an entry -- not OK
				err = io.ErrUnexpectedEOF
				return
			}
		}

		lines = append(lines, r.scanner.Text())
	}

	return Entry{Lines: lines}, nil
}
