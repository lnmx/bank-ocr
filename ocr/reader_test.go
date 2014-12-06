package ocr

import (
	"io"
	"strings"
	"testing"
)

func TestReaderOk(t *testing.T) {
	reader := NewEntryReader(4, strings.NewReader("1\n2\n3\n4\n"))

	entry, err := reader.Next()

	if err != nil {
		t.Fatal(err)
	}

	if entry == nil {
		t.Fatal("no entry returned")
	}

	if len(entry.Lines) != 4 {
		t.Fatal("expected 4 lines in parsed entry")
	}

	if entry.Lines[0] != "1" || entry.Lines[3] != "4" {
		t.Fatal("wrong lines in parsed entry")
	}
}

func TestReaderEOF(t *testing.T) {
	reader := NewEntryReader(4, strings.NewReader("1\n2\n3"))

	entry, err := reader.Next()

	if entry != nil {
		t.Fatal("expected nil entry from short input")
	}

	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected", io.ErrUnexpectedEOF, "got", err)
	}
}
