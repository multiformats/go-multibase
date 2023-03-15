package multibase

import (
	"bytes"
	"crypto/rand"
	"sort"
	"testing"
)

func TestMap(t *testing.T) {
	for s, e := range Encodings {
		s2 := EncodingToStr[e]
		if s != s2 {
			t.Errorf("round trip failed on encoding map: %s != %s", s, s2)
		}
	}
	for e, s := range EncodingToStr {
		e2 := Encodings[s]
		if e != e2 {
			t.Errorf("round trip failed on encoding map: '%c' != '%c'", e, e2)
		}
	}
}

var sampleBytes = []byte("Decentralize everything!!!")
var encodedSamples = map[Encoding]string{
	Identity:          string(rune(0x00)) + "Decentralize everything!!!",
	Base2:             "00100010001100101011000110110010101101110011101000111001001100001011011000110100101111010011001010010000001100101011101100110010101110010011110010111010001101000011010010110111001100111001000010010000100100001",
	Base16:            "f446563656e7472616c697a652065766572797468696e67212121",
	Base16Upper:       "F446563656E7472616C697A652065766572797468696E67212121",
	Base32:            "birswgzloorzgc3djpjssazlwmvzhs5dinfxgoijbee",
	Base32Upper:       "BIRSWGZLOORZGC3DJPJSSAZLWMVZHS5DINFXGOIJBEE",
	Base32pad:         "cirswgzloorzgc3djpjssazlwmvzhs5dinfxgoijbee======",
	Base32padUpper:    "CIRSWGZLOORZGC3DJPJSSAZLWMVZHS5DINFXGOIJBEE======",
	Base32hex:         "v8him6pbeehp62r39f9ii0pbmclp7it38d5n6e89144",
	Base32hexUpper:    "V8HIM6PBEEHP62R39F9II0PBMCLP7IT38D5N6E89144",
	Base32hexPad:      "t8him6pbeehp62r39f9ii0pbmclp7it38d5n6e89144======",
	Base32hexPadUpper: "T8HIM6PBEEHP62R39F9II0PBMCLP7IT38D5N6E89144======",
	Base36:            "km552ng4dabi4neu1oo8l4i5mndwmpc3mkukwtxy9",
	Base36Upper:       "KM552NG4DABI4NEU1OO8L4I5MNDWMPC3MKUKWTXY9",
	Base58BTC:         "z36UQrhJq9fNDS7DiAHM9YXqDHMPfr4EMArvt",
	Base58Flickr:      "Z36tpRGiQ9Endr7dHahm9xwQdhmoER4emaRVT",
	Base64:            "mRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE",
	Base64url:         "uRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE",
	Base64pad:         "MRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE=",
	Base64urlPad:      "URGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE=",
	Base256Emoji:      "🚀💛✋💃✋😻😈🥺🤤🍀🌟💐✋😅✋💦✋🥺🏃😈😴🌟😻😝👏👏👏",
}

func testEncode(t *testing.T, encoding Encoding, bytes []byte, expected string) {
	actual, err := Encode(encoding, bytes)
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("encoding failed for %c (%d / %s), expected: %s, got: %s", encoding, encoding, EncodingToStr[encoding], expected, actual)
	}
}

func testDecode(t *testing.T, expectedEncoding Encoding, expectedBytes []byte, data string) {
	actualEncoding, actualBytes, err := Decode(data)
	if err != nil {
		t.Error(err)
		return
	}
	if actualEncoding != expectedEncoding {
		t.Errorf("wrong encoding code, expected: %c (%d), got %c (%d)", expectedEncoding, expectedEncoding, actualEncoding, actualEncoding)
	}
	if !bytes.Equal(actualBytes, expectedBytes) {
		t.Errorf("decoding failed for %c (%d), expected: %v, got %v", actualEncoding, actualEncoding, expectedBytes, actualBytes)
	}
}

func TestEncode(t *testing.T) {
	for encoding := range EncodingToStr {
		testEncode(t, encoding, sampleBytes, encodedSamples[encoding])
	}
}

func TestDecode(t *testing.T) {
	for encoding := range EncodingToStr {
		testDecode(t, encoding, sampleBytes, encodedSamples[encoding])
	}
}

func TestRoundTrip(t *testing.T) {

	for base := range EncodingToStr {
		if int(base) == 0 {
			// skip identity: any byte goes there
			continue
		}

		_, _, err := Decode(string(rune(base)) + "\u00A0")
		if err == nil {
			t.Fatal(EncodingToStr[base] + " decode should fail on low-unicode")
		}

		_, _, err = Decode(string(rune(base)) + "\u1F4A8")
		if err == nil {
			t.Fatal(EncodingToStr[base] + " decode should fail on emoji")
		}

		_, _, err = Decode(string(rune(base)) + "!")
		if err == nil {
			t.Fatal(EncodingToStr[base] + " decode should fail on punctuation")
		}

		_, _, err = Decode(string(rune(base)) + "\xA0")
		if err == nil {
			t.Fatal(EncodingToStr[base] + " decode should fail on high-latin1")
		}
	}

	buf := make([]byte, 137+16) // sufficiently large prime number of bytes + another 16 to test leading 0s
	rand.Read(buf[16:])

	for base := range EncodingToStr {

		// test roundtrip from the full zero-prefixed buffer down to a single byte
		for i := 0; i < len(buf); i++ {

			// use a copy to verify we are not overwriting the supplied buffer
			newBuf := make([]byte, len(buf)-i)
			copy(newBuf, buf[i:])

			enc, err := Encode(base, newBuf)
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

			if !bytes.Equal(newBuf, buf[i:]) {
				t.Fatal("the provided buffer was modified", buf[i:], out)
			}

			if !bytes.Equal(buf[i:], out) {
				t.Fatal("input wasnt the same as output", buf[i:], out)
			}

			// When we have 3 leading zeroes, do a few extra tests
			// ( choice of leading zeroes is arbitrary - just cutting down on test permutations )

			if i == 13 {

				// if this is a case-insensitive codec semi-randomly swap case in enc and try again
				name := EncodingToStr[base]
				if name[len(name)-5:] == "upper" || Encodings[name+"upper"] > 0 {
					caseTamperedEnc := []byte(enc)

					for _, j := range []int{3, 5, 8, 13, 21, 23, 29, 47, 52} {
						if caseTamperedEnc[j] >= 65 && caseTamperedEnc[j] <= 90 {
							caseTamperedEnc[j] += 32
						} else if caseTamperedEnc[j] >= 97 && caseTamperedEnc[j] <= 122 {
							caseTamperedEnc[j] -= 32
						}
					}

					e, out, err := Decode(string(caseTamperedEnc))
					if err != nil {
						t.Fatal(err)
					}

					if e != base {
						t.Fatal("got wrong encoding out")
					}
					if !bytes.Equal(buf[i:], out) {
						t.Fatal("input wasn't the same as output", buf[i:], out)
					}
				}
			}
		}
	}

	// Test that nothing overflows
	maxValueBuf := make([]byte, 131)
	for i := 0; i < len(maxValueBuf); i++ {
		maxValueBuf[i] = 0xFF
	}

	for base := range EncodingToStr {

		// test roundtrip from the complete buffer down to a single byte
		for i := 0; i < len(maxValueBuf); i++ {

			enc, err := Encode(base, maxValueBuf[i:])
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

			if !bytes.Equal(maxValueBuf[i:], out) {
				t.Fatal("input wasn't the same as output", maxValueBuf[i:], out)
			}
		}
	}

	_, _, err := Decode("")
	if err == nil {
		t.Fatal("shouldn't be able to decode empty string")
	}
}

var benchmarkBuf [36]byte // typical CID size
var benchmarkCodecs []string

func init() {
	rand.Read(benchmarkBuf[:])

	benchmarkCodecs = make([]string, 0, len(Encodings))
	for n := range Encodings {

		// // Only bench b36 and b58
		// if len(n) < 6 || (n[4:6] != "36" && n[4:6] != "58") {
		// 	continue
		// }

		benchmarkCodecs = append(benchmarkCodecs, n)
	}
	sort.Strings(benchmarkCodecs)
}

func BenchmarkRoundTrip(b *testing.B) {
	b.ResetTimer()

	for _, name := range benchmarkCodecs {
		b.Run(name, func(b *testing.B) {
			base := Encodings[name]
			for i := 0; i < b.N; i++ {
				enc, err := Encode(base, benchmarkBuf[:])
				if err != nil {
					b.Fatal(err)
				}

				e, out, err := Decode(enc)
				if err != nil {
					b.Fatal(err)
				}

				if e != base {
					b.Fatal("got wrong encoding out")
				}

				if !bytes.Equal(benchmarkBuf[:], out) {
					b.Fatal("input wasnt the same as output", benchmarkBuf, out)
				}
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	b.ResetTimer()

	for _, name := range benchmarkCodecs {
		b.Run(name, func(b *testing.B) {
			base := Encodings[name]
			for i := 0; i < b.N; i++ {
				_, err := Encode(base, benchmarkBuf[:])
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ResetTimer()

	for _, name := range benchmarkCodecs {
		b.Run(name, func(b *testing.B) {
			enc, _ := Encode(Encodings[name], benchmarkBuf[:])
			for i := 0; i < b.N; i++ {
				_, _, err := Decode(enc)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
