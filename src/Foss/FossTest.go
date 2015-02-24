package Foss

import (
	"os.File"
	"testing"
)

func TestMode(t *testing.T) {
	file, err := File.Create("test.txt")
	if err != nil {
		testing.Fatal(err)
	} else {

	}
}
