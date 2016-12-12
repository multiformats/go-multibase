package multibase

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	buf := make([]byte, 17)
	rand.Read(buf)

	baseList := []int{ Base16, Base32, Base32hex, Base32pad, Base32hexPad, Base58BTC, Base58Flickr, Base64pad, Base64urlPad, Identity }

	for _, base := range baseList {
		enc, err := Encode(base, buf)
		if err != nil {
			t.Fatal(err)
		}

		e, out, err := Decode(enc)
		if err != nil {
			t.Fatal(err)
		}

		if e != base {
			t.Fatal("got wrong encoding out")
		}

		if !bytes.Equal(buf, out) {
			t.Fatal("input wasnt the same as output", buf, out)
		}
	}

	_, _, err := Decode("")
	if err == nil {
		t.Fatal("shouldnt be able to decode empty string")
	}
}
