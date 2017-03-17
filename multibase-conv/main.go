package main

import (
	"fmt"
	"os"

	multibase "github.com/multiformats/go-multibase"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s NEW-BASE CID...\n", os.Args[0])
		os.Exit(1)
	}

	newBase := os.Args[1]
	cids := os.Args[2:]

	for _, cid := range cids {
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
		fmt.Println(newCid)
	}

}
