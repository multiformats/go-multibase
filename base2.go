package multibase

import (
	"fmt"
	"math"
	"strconv"
)

// binaryEncodeToString takes an array of bytes and returns
// multibase binary representation
func binaryEncodeToString(src []byte) string {
	dst := make([]byte, len(src)*8)
	encodeBinary(dst, src)
	return string(dst)
}

// encodeBinary takes the src and dst bytes and converts each
// byte to their binary rep using power reduction method
func encodeBinary(dst []byte, src []byte) {
	for i := 0; i < len(src); i++ {
		t := src[i]
		for j := i * 8; j < (i*8)+8; j++ {
			higherPower := math.Pow(2, float64(7-(j%8)))
			if float64(t) >= higherPower {
				dst[j] = '1'
				t = t - byte(higherPower)
			} else {
				dst[j] = '0'
			}
		}
	}
}

// decodeBinaryString takes multibase binary representation
// and returns a byte array
func decodeBinaryString(s string) ([]byte, error) {
	if len(s)%8 != 0 {
		return nil, fmt.Errorf("cannot decode multibase: %s",
			"length should be a multiple of 4")
	}

	data := make([]byte, len(s)/8)

	for i, dstIndex := 0, 0; i < len(s); i = i + 8 {
		value, err := strconv.ParseInt(s[i:i+8], 2, 0)
		if err != nil {
			return nil, fmt.Errorf("error while conversion: %s", err)
		}

		data[dstIndex] = byte(value)
		dstIndex++
	}

	return data, nil
}
