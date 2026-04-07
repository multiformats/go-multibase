package multibase

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestSpec(t *testing.T) {
	file, err := os.Open("spec/multibase.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = false
	reader.FieldsPerRecord = 5
	reader.TrimLeadingSpace = true

	values, err := reader.ReadAll()
	if err != nil {
		t.Error(err)
	}
	specEncodings := make(map[Encoding]string, len(values)-1)
	for _, v := range values[1:] {
		unicodeStr := v[0] // e.g. "U+007A"
		encoding := v[2]

		var code Encoding
		if !strings.HasPrefix(unicodeStr, "U+") {
			t.Errorf("unexpected unicode format %q", unicodeStr)
			continue
		}
		i, err := strconv.ParseUint(unicodeStr[2:], 16, 64)
		if err != nil {
			t.Errorf("invalid unicode codepoint %q", unicodeStr)
			continue
		}
		code = Encoding(i)
		specEncodings[code] = encoding
	}

	for name, enc := range Encodings {
		specName, ok := specEncodings[enc]
		if !ok {
			t.Errorf("encoding %q (%c) not defined in the spec", name, enc)
			continue
		}
		// The spec marks some code points as "none"/reserved that
		// go-multibase implements (e.g. identity at 0x00).
		if specName == "none" {
			continue
		}
		if specName != name {
			t.Errorf("encoding %q (%c) has unexpected name %q", specName, enc, name)
		}
	}
}
func TestSpecVectors(t *testing.T) {
	files, err := filepath.Glob("spec/tests/*.csv")
	if err != nil {
		t.Fatal(err)
	}
	for _, fname := range files {
		t.Run(fname, func(t *testing.T) {
			file, err := os.Open(fname)
			if err != nil {
				t.Error(err)
				return
			}
			defer file.Close()
			reader := csv.NewReader(file)
			reader.LazyQuotes = false
			reader.FieldsPerRecord = 2
			reader.TrimLeadingSpace = true

			values, err := reader.ReadAll()
			if err != nil {
				t.Error(err)
			}
			if len(values) == 0 {
				t.Error("no test values")
				return
			}
			header := values[0]

			var decodeOnly bool
			switch header[0] {
			case "encoding":
			case "non-canonical encoding":
				decodeOnly = true
			default:
				t.Errorf("invalid test spec %q", fname)
				return
			}

			testValue, err := strconv.Unquote("\"" + header[1] + "\"")
			if err != nil {
				t.Error("failed to unquote testcase:", err)
				return
			}

			for _, testCase := range values[1:] {
				encodingName := testCase[0]
				expected := testCase[1]

				t.Run(encodingName, func(t *testing.T) {
					encoder, err := EncoderByName(encodingName)
					if err != nil {
						t.Skipf("skipping %s: not supported", encodingName)
						return
					}
					if !decodeOnly {
						t.Logf("encoding %q with %s", testValue, encodingName)
						actual := encoder.Encode([]byte(testValue))
						if expected != actual {
							t.Errorf("expected %q, got %q", expected, actual)
						}
					}
					t.Logf("decoding %q", expected)
					encoding, decoded, err := Decode(expected)
					if err != nil {
						t.Error("failed to decode:", err)
						return
					}
					expectedEncoding := Encodings[encodingName]
					if encoding != expectedEncoding {
						t.Errorf("expected encoding to be %c, got %c", expectedEncoding, encoding)
					}
					if string(decoded) != testValue {
						t.Errorf("failed to decode %q to %q, got %q", expected, testValue, string(decoded))
					}
				})

			}
		})
	}
}

func FuzzDecode(f *testing.F) {
	files, err := filepath.Glob("spec/tests/*.csv")
	if err != nil {
		f.Fatal(err)
	}
	for _, fname := range files {
		func() {
			file, err := os.Open(fname)
			if err != nil {
				f.Fatal(err)
			}
			defer file.Close()
			reader := csv.NewReader(file)
			reader.LazyQuotes = false
			reader.FieldsPerRecord = 2
			reader.TrimLeadingSpace = true

			values, err := reader.ReadAll()
			if err != nil {
				f.Fatal(err)
			}

			for _, tc := range values[1:] {
				f.Add(tc[1])
			}
		}()
	}

	f.Fuzz(func(_ *testing.T, data string) {
		Decode(data)
	})
}
