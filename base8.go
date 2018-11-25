package multibase

import (
	"fmt"
	"strconv"
)

func octalEncodeToString(src []byte) string {
	dst := make([]byte, len(src)*3)
	octalEncode(dst, src)
	return string(dst)
}

var octalTable = [8]byte{
	'0', '1', '2', '3', '4', '5', '6', '7',
}

func octalEncode(dst, src []byte) int {
	for i, v := range src {
		dst[i*3+2] = octalTable[v&0x7]
		v = v >> 3
		dst[i*3+1] = octalTable[v&0x7]
		v = v >> 3
		dst[i*3] = octalTable[v&0x7]
		v = v >> 3
	}

	return len(src) * 3
}

// decodeOctalString takes an octal string and gives byte array of decimals
func decodeOctalString(s string) ([]byte, error) {
	data := make([]byte, len(s)/3)
	if len(s)%3 != 0 {
		return nil, fmt.Errorf("cannot decode multibase: %s",
			"length should be a multiple of 3")
	}

	for i, dstIndex := 0, 0; i < len(s); i = i + 3 {
		value, err := strconv.ParseInt(s[i:i+3], 8, 8)
		if err != nil {
			return nil, fmt.Errorf("error while conversion: %s", err)
		}

		data[dstIndex] = byte(value)
		dstIndex++
	}

	return data, nil
}
