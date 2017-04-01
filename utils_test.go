package goutils

import (
	"testing"
)

func TestFileAndExt(t *testing.T) {
	type fileAndExt struct {
		file string
		ext  string
	}
	type test struct {
		input    string
		expected fileAndExt
	}
	data := []test{
		{"test.txt", fileAndExt{"test", ".txt"}},
		{"path/test.txt", fileAndExt{"test", ".txt"}},
		{"../relativepath/test.txt", fileAndExt{"test", ".txt"}},
		{"/absolutepath/test.txt", fileAndExt{"test", ".txt"}},
	}

	for i, d := range data {
		file, ext := FileAndExt(d.input)
		if file != d.expected.file {
			t.Errorf("Test %d failed. Expected %s got %s", i, d.expected.file, file)
		}
		if ext != d.expected.ext {
			t.Errorf("Test %d failed. Expected %s got %s", i, d.expected.ext, ext)
		}
	}
}
