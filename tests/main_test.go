package test

import (
	"github.com/Harmohit-Singh/punjabi_dictionary/src"
	"os"
	"testing"
)

func TestOpenPageNumber(t *testing.T) {
	expected, _ := os.ReadFile(os.Getenv("HOME") + "/Desktop/punjabi_dictionary/sample_data/page201.txt")
	result := src.Open_page_number(201)
	if result != string(expected) {
		t.Errorf("open_page_number(valid_page_number) was incorrect, got: %s, want: %s", result, expected)
	}
}
