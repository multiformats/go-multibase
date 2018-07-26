package multibase

import (
	"fmt"
)

// Prefix is a multibase encoding that is verified to be supported and
// supports an Encode method that does not return an error
type Prefix struct {
	enc Encoding
}

// NewPrefix create a new Prefix type from either an Encoding or a
// string.  If nil or an empty string the default base (currently
// Base58BTC) will be used.
func NewPrefix(e interface{}) (Prefix, error) {
	base := Encoding(Base58BTC)
	switch v := e.(type) {
	case byte:
		base = Encoding(v)
	case rune:
		base = Encoding(v)
	case int:
		base = Encoding(v)
	case Encoding:
		base = v
	case string:
		ok := true
		// note: if empty string use default value
		if len(v) == 1 {
			base = Encoding(v[0])
			_, ok = EncodingToStr[base]
		} else if len(v) > 1 {
			base, ok = Encodings[v]
		}
		if !ok {
			return Prefix{-1}, fmt.Errorf("Unsupported multibase encoding: %s", v)
		}
		return Prefix{base}, nil
	case nil:
		/* use default value */
	default:
		return Prefix{-1}, fmt.Errorf("Unsupported parameter type.")
	}
	_, ok := EncodingToStr[base]
	if !ok {
		return Prefix{-1}, fmt.Errorf("Unsupported multibase encoding: %d", base)
	}
	return Prefix{base}, nil
}

func (p Prefix) Encoding() Encoding {
	return p.enc
}

// Encode encodes the multibase using the given Prefix.  
func (p Prefix) Encode(data []byte) string {
	str, err := Encode(p.enc, data)
	if err != nil {
		// should not happen
		panic(err)
	}
	return str
}
