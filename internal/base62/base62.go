package base62

import (
	"errors"
	"strings"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	var sb strings.Builder
	for num > 0 {
		remainder := num % 62
		sb.WriteByte(alphabet[remainder])
		num /= 62
	}

	str := sb.String()
	reversed := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		reversed[i] = str[len(str)-1-i]
	}

	return string(reversed)
}

// Returns an error if the input contains invalid characters or overflows uint64.
func Decode(encoded string) (uint64, error) {
	if encoded == "" {
		return 0, errors.New("cannot decode empty string")
	}

	var result uint64 = 0
	for i := 0; i < len(encoded); i++ {
		char := encoded[i]
		
		var value int
		switch {
		case char >= '0' && char <= '9':
			value = int(char - '0')
		case char >= 'A' && char <= 'Z':
			value = int(char - 'A') + 10
		case char >= 'a' && char <= 'z':
			value = int(char - 'a') + 36
		default:
			return 0, errors.New("invalid character in base62 string: " + string(char))
		}

		// Prevent multiplication overflow bounds
		if result > (18446744073709551615-uint64(value))/62 {
			return 0, errors.New("base62 value overflows uint64")
		}

		result = result*62 + uint64(value)
	}

	return result, nil
}