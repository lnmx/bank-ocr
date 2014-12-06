package ocr

import (
	"strings"
)

type Entry struct {
	Lines []string
}

func (e Entry) String() string {
	return strings.Join(e.Lines, "\n")
}
