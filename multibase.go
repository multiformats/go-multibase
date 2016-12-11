package multibase

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"

	b58 "github.com/jbenet/go-base58"
	b32 "github.com/whyrusleeping/base32"
)

const (
	Base1        = '1'
	Base2        = '0'
	Base8        = '7'
	Base10       = '9'
	Base16       = 'f' // 'F'
	Base32       = 'u' // 'U'
	Base32hex    = 'v' // 'V'
	Base58Flickr = 'Z'
	Base58BTC    = 'z'
	Base64       = 'y'
	Base64url    = 'Y'
	Binary       = 'X'
)

var ErrUnsupportedEncoding = fmt.Errorf("selected encoding not supported")

func Encode(base int, data []byte) (string, error) {
	switch base {
	case Base16, 'F':
		return string(Base16) + hex.EncodeToString(data), nil
	case Base32, 'U':
		return string(Base32) + b32.RawStdEncoding.EncodeToString(data), nil
	case Base32hex, 'V':
		return string(Base32hex) + b32.RawHexEncoding.EncodeToString(data), nil
	case Base58BTC:
		return string(Base58BTC) + b58.EncodeAlphabet(data, b58.BTCAlphabet), nil
	case Base58Flickr:
		return string(Base58Flickr) + b58.EncodeAlphabet(data, b58.FlickrAlphabet), nil
	case Base64:
		return string(Base64) + base64.StdEncoding.EncodeToString(data), nil
	case Base64url:
		return string(Base64url) + base64.URLEncoding.EncodeToString(data), nil
	case Binary:
		return string(Binary) + string(data), nil
	default:
		return "", ErrUnsupportedEncoding
	}
}

func Decode(data string) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, fmt.Errorf("cannot decode multibase for zero length string")
	}

	switch data[0] {
	case Base16, 'F':
		bytes, err := hex.DecodeString(data[1:])
		return Base16, bytes, err
	case Base32, 'U':
		bytes, err := b32.RawStdEncoding.DecodeString(data[1:])
		return Base32, bytes, err
	case Base32hex, 'V':
		bytes, err := b32.RawHexEncoding.DecodeString(data[1:])
		return Base32hex, bytes, err
	case Base58BTC:
		return Base58BTC, b58.DecodeAlphabet(data[1:], b58.BTCAlphabet), nil
	case Base58Flickr:
		return Base58Flickr, b58.DecodeAlphabet(data[1:], b58.FlickrAlphabet), nil
	case Base64:
		bytes, err := base64.StdEncoding.DecodeString(data[1:])
		return Base64, bytes, err
	case Base64url:
		bytes, err := base64.URLEncoding.DecodeString(data[1:])
		return Base64url, bytes, err
	case Binary:
		return Binary, []byte(data[1:]), nil
	default:
		return -1, nil, ErrUnsupportedEncoding
	}
}
