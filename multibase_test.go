package multibase

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestBase58RoundTrip(t *testing.T) {
	buf := make([]byte, 16)
	rand.Read(buf)

	enc, err := Encode(Base58BTC, buf)
	if err != nil {
		t.Fatal(err)
	}

	e, out, err := Decode(enc)
	if err != nil {
		t.Fatal(err)
	}

	if e != Base58BTC {
		t.Fatal("got wrong encoding out")
	}

	if !bytes.Equal(buf, out) {
		t.Fatal("input wasnt the same as output", buf, out)
	}

	_, _, err = Decode("")
	if err == nil {
		t.Fatal("shouldnt be able to decode empty string")
	}
}
