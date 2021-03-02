package main

import (
	"flag"
	"io/ioutil"

	"github.com/anierzad/cr2-raw-go/read"
)

func main() {

	// Get file path from passed parameters.
	filePath := flag.String("f", "", "Path to a .CR2 file.")
	flag.Parse()

	// Load data from file.
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	// Create a new tiff head reader, passing a pointer to the data.
	thr := read.NewTiffHeadReader(&data)

	_ = thr
}
