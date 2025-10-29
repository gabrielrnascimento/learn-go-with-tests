package main

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := tape{file}

	_, _ = tape.Write([]byte("abc"))

	_, _ = file.Seek(0, io.SeekStart)
	newFileConstants, _ := io.ReadAll(file)

	got := string(newFileConstants)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
