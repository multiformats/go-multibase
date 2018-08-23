package multibase

import (
	"testing"
)

func TestInvalidPrefix(t *testing.T) {
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

func TestPrefix(t *testing.T) {
	for str, base := range Encodings {
		prefix, err := NewEncoder(base)
		if err != nil {
			t.Fatalf("NewEncoder(%c) failed: %v", base, err)
		}
		str1, err := Encode(base, sampleBytes)
		if err != nil {
			t.Fatal(err)
		}
		str2 := prefix.Encode(sampleBytes)
		if str1 != str2 {
			t.Errorf("encoded string mismatch: %s != %s", str1, str2)
		}
		_, err = EncoderByName(str)
		if err != nil {
			t.Fatalf("NewEncoder(%s) failed: %v", str, err)
		}
	}
}
