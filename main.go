package main

import (
	"flag"
	"fmt"
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

	// Create a new idf reader, passing offset and pointer to the data.
	ir := read.NewIfdReader(thr.FirstIfdOffset(), &data)
	irNo := 0

	for {

		// Print details about ifd.
		entries, _ := ir.GetIfdEntries()

		for k, e := range entries {
			fmt.Printf("%v: %v\n", k, e)
		}

		// Check there is another ifd to read.
		if ir.NextIfdOffset() == 0 {
			break
		}

		// Move to next ifd.
		ir = read.NewIfdReader(ir.NextIfdOffset(), &data)
		irNo++
	}
}
