package multibase

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"

	b58 "github.com/jbenet/go-base58"
	b32 "github.com/whyrusleeping/base32"
)

const (
	Identity          = 0x00
	Base1             = '1'
	Base2             = '0'
	Base8             = '7'
	Base10            = '9'
	Base16            = 'f'
	Base16Upper       = 'F'
	Base32            = 'b'
	Base32Upper       = 'B'
	Base32pad         = 'c'
	Base32padUpper    = 'C'
	Base32hex         = 'v'
	Base32hexUpper    = 'V'
	Base32hexPad      = 't'
	Base32hexPadUpper = 'T'
	Base58Flickr      = 'Z'
	Base58BTC         = 'z'
	Base64            = 'm'
	Base64url         = 'u'
	Base64pad         = 'M'
	Base64urlPad      = 'U'
)

var ErrUnsupportedEncoding = fmt.Errorf("selected encoding not supported")

func Encode(base int, data []byte) (string, error) {
	switch base {
	case Identity:
		return string(Identity) + string(data), nil
	case Base16, Base16Upper:
		return string(Base16) + hex.EncodeToString(data), nil
	case Base32, Base32Upper:
		return string(Base32) + b32.RawStdEncoding.EncodeToString(data), nil
	case Base32hex, Base32hexUpper:
		return string(Base32hex) + b32.RawHexEncoding.EncodeToString(data), nil
	case Base32pad, Base32padUpper:
		return string(Base32pad) + b32.StdEncoding.EncodeToString(data), nil
	case Base32hexPad, Base32hexPadUpper:
		return string(Base32hexPad) + b32.HexEncoding.EncodeToString(data), nil
	case Base58BTC:
		return string(Base58BTC) + b58.EncodeAlphabet(data, b58.BTCAlphabet), nil
	case Base58Flickr:
		return string(Base58Flickr) + b58.EncodeAlphabet(data, b58.FlickrAlphabet), nil
	case Base64pad:
		return string(Base64pad) + base64.StdEncoding.EncodeToString(data), nil
	case Base64urlPad:
		return string(Base64urlPad) + base64.URLEncoding.EncodeToString(data), nil
	default:
		return "", ErrUnsupportedEncoding
	}
}

func Decode(data string) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, fmt.Errorf("cannot decode multibase for zero length string")
	}

	switch data[0] {
	case Identity:
		return Identity, []byte(data[1:]), nil
	case Base16, Base16Upper:
		bytes, err := hex.DecodeString(data[1:])
		return Base16, bytes, err
	case Base32, Base32Upper:
		bytes, err := b32.RawStdEncoding.DecodeString(data[1:])
		return Base32, bytes, err
	case Base32hex, Base32hexUpper:
		bytes, err := b32.RawHexEncoding.DecodeString(data[1:])
		return Base32hex, bytes, err
	case Base32pad, Base32padUpper:
		bytes, err := b32.StdEncoding.DecodeString(data[1:])
		return Base32pad, bytes, err
	case Base32hexPad, Base32hexPadUpper:
		bytes, err := b32.HexEncoding.DecodeString(data[1:])
		return Base32hexPad, bytes, err
	case Base58BTC:
		return Base58BTC, b58.DecodeAlphabet(data[1:], b58.BTCAlphabet), nil
	case Base58Flickr:
		return Base58Flickr, b58.DecodeAlphabet(data[1:], b58.FlickrAlphabet), nil
	case Base64pad:
		bytes, err := base64.StdEncoding.DecodeString(data[1:])
		return Base64pad, bytes, err
	case Base64urlPad:
		bytes, err := base64.URLEncoding.DecodeString(data[1:])
		return Base64urlPad, bytes, err
	default:
		return -1, nil, ErrUnsupportedEncoding
	}
}
