package main

import (
	"fmt"
	"os"
	
	multibase "github.com/multiformats/go-multibase"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s CID NEW-BASE\n", os.Args[0])
		os.Exit(1)
	}

	cid := os.Args[1]
	newBase := os.Args[2]

	_, data, err := multibase.Decode(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	newCid, err := multibase.Encode(multibase.Encoding(newBase[0]), data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", newCid)
}
