package multibase

import (
	"testing"
	"unicode/utf8"
)

func TestInvalidCode(t *testing.T) {
	_, err := NewEncoder('q')
	if err == nil {
		t.Error("expected failure")
	}
}

func TestInvalidName(t *testing.T) {
	values := []string{"invalid", "", "q"}
	for _, val := range values {
		_, err := EncoderByName(val)
		if err == nil {
			t.Errorf("EncoderByName(%v) expected failure", val)
		}
	}
}

func TestEncoder(t *testing.T) {
	for name, code := range Encodings {
		encoder, err := NewEncoder(code)
		if err != nil {
			t.Fatal(err)
		}
		// Make sure the MustNewEncoder doesn't panic
		MustNewEncoder(code)
		str, err := Encode(code, sampleBytes)
		if err != nil {
			t.Fatal(err)
		}
		str2 := encoder.Encode(sampleBytes)
		if str != str2 {
			t.Errorf("encoded string mismatch: %s != %s", str, str2)
		}
		_, err = EncoderByName(name)
		if err != nil {
			t.Fatalf("EncoderByName(%s) failed: %v", name, err)
		}
		// Test that an encoder can be created from the single letter
		// prefix
		r, _ := utf8.DecodeRuneInString(str)
		_, err = EncoderByName(string(r))
		if err != nil {
			t.Fatalf("EncoderByName(%s) failed: %v", string(r), err)
		}
	}
}
