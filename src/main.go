package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get current directory and print
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Sys())
	}

}

// 1. Print file mode - fileinfo.Mode() - prints mode

// &{main.go 297 420 {63560235597 0 0x16e1e0} 0x2081da120}

// type FileInfo interface {
//         Name() string       // base name of the file
//         Size() int64        // length in bytes for regular files; system-dependent for others
//         Mode() FileMode     // file mode bits
//         ModTime() time.Time // modification time
//         IsDir() bool        // abbreviation for Mode().IsDir()
//         Sys() interface{}   // underlying data source (can return nil)
// }
