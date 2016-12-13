package multibase

import (
	"bytes"
	"math/rand"
	"testing"
)

var sampleBytes = []byte("Decentralize everything!!")
var encodedSamples = map[int]string{
	Identity: string(0x00) + "Decentralize everything!!",
	Base64pad: "MRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchIQ==",
}

func testEncode(t *testing.T, encoding int, bytes []byte, expected string)  {
	actual, err := Encode(encoding, bytes)
	if err != nil {
		t.Error(err)
		return
	}
	assertEqual(t, expected, actual, "Encoding failure for encoding %c (%d)", encoding, encoding)
}

func testDecode(t *testing.T, expectedEncoding int, expectedBytes []byte, data string)  {
	actualEncoding, actualBytes, err := Decode(data)
	if err != nil {
		t.Error(err)
		return
	}
	assertEqual(t, expectedEncoding, actualEncoding)
	assertEqual(t, expectedBytes, actualBytes, "Encoding failure for encoding %c (%d)", expectedEncoding, expectedEncoding)
}

func TestEncode(t *testing.T)  {
	for encoding, data := range encodedSamples {
		testEncode(t, encoding, sampleBytes, data)
	}
}

func TestDecode(t *testing.T)  {
	for encoding, data := range encodedSamples {
		testDecode(t, encoding, sampleBytes, data)
	}
}

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
