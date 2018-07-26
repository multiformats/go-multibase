package multibase

import (
	"testing"
)

func TestInvalidPrefix(t *testing.T) {
	_, err := NewPrefix('q')
	if err == nil {
		t.Error("expected failure")
	}
}

func TestPrefix(t *testing.T) {
	prefix,err := NewPrefix(nil)
	if err != nil {
		t.Fatal(err)
	}
	if prefix.Encoding() != Base58BTC {
		t.Error("unexpected default encoding")
	}
	for str, base := range Encodings {
		prefix,err = NewPrefix(base)
		if err != nil {
			t.Fatalf("NewPrefix(%c) failed: %v", base, err)
		}
		str1, err := Encode(base, sampleBytes)
		if err != nil {
			t.Fatal(err)
		}
		str2 := prefix.Encode(sampleBytes)
		if str1 != str2 {
			t.Errorf("encoded string mismatch: %s != %s", str1, str2)
		}
		_, err = NewPrefix(str)
		if err != nil {
			t.Fatalf("NewPrefix(%s) failed: %v", str, err)
		}
	}
}
