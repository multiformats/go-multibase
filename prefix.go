package multibase

import ()

type Prefix struct {
	enc Encoding
}

func NewPrefix(e Encoding) (Prefix, error) {
	_, err := Encode(e, nil)
	if err != nil {
		return Prefix{-1}, err
	}
	return Prefix{e}, nil
}

func (p Prefix) Encoding() Encoding {
	return p.enc
}

func (p Prefix) Encode(data []byte) string {
	str, err := Encode(p.enc, data)
	if err != nil {
		// should not happen
		panic(err)
	}
	return str
}
