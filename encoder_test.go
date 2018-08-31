package multibase

import (
	"testing"
)

func TestInvalidEncoding(t *testing.T) {
	err := CheckEncoding(Encoding('q'))
	if err == nil {
		t.Errorf("CheckEncoding('q') expected failure")
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
		encoder := NewEncoder(code)
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
			t.Errorf("EncoderByName(%s) failed: %v", name, err)
		}
		// Test that an encoder can be created from the single letter
		// prefix
		_, err = EncoderByName(str[0:1])
		if err != nil {
			t.Errorf("EncoderByName(%s) failed: %v", str[0:1], err)
		}
	}
}
