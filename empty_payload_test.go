package multibase

import "testing"

// Encoding an empty byte slice yields a string containing only the encoding
// prefix. Decode must accept that form for every supported encoding.

func TestEmptyPayloadRoundTrip(t *testing.T) {
	for name, enc := range Encodings {
		s, err := Encode(enc, []byte{})
		if err != nil {
			t.Fatalf("%s: Encode error: %v", name, err)
		}

		gotEnc, b, err := Decode(s)
		if err != nil {
			t.Errorf("%s: Decode(%q) error: %v", name, s, err)
			continue
		}
		if gotEnc != enc {
			t.Errorf("%s: Decode(%q) encoding = %v, want %v", name, s, gotEnc, enc)
		}
		if len(b) != 0 {
			t.Errorf("%s: Decode(%q) = %x, want empty", name, s, b)
		}
	}
}
