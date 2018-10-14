package main

import (
	"fmt"
)

const (
	// digitSet62 is the set of 62 digits we chose to represent numbers in base-62.
	digitSet62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// maxLen is the number of digits in "lYGhA16ahyf"
	// base-62(math.MaxUint64) = "lYGhA16ahyf"
	maxLen = 11
)

// digitValue returns the integer value of base-62 digit d.
// ok is false if d is not a valid digit.
func digitValue(d byte) (val uint64, ok bool) {
	for i := 0; i < 62; i++ {
		if digitSet62[i] == d {
			return uint64(i), true
		}
	}
	return
}

// toBase62 converts v into its base-62 representation.
func toBase62(v uint64) string {
	// create a buffer to store 11 base-62 digits
	buf := make([]byte, 0)
	// 0-padding
	for len(buf) < 11 {
		buf = append(buf, '0')
	}

	// convert id to a base 62 number
	i := 10
	for v > 0 {
		buf[i] = digitSet62[v%62]
		v = v / 62
		i--
	}

	return string(buf)
}

// fromBase62 converts base-62 digits into an uint64.
func fromBase62(digits string) (uint64, error) {
	var id uint64

	if digits == "" {
		return 0, fmt.Errorf("fromBase62: can't convert empty string")
	}

	for c := 0; c < len(digits); c++ {
		i, ok := digitValue(digits[c])
		if !ok {
			return 0, fmt.Errorf("fromBase62: unexpected char %c ", digits[c])
		}
		id = id*62 + uint64(i)
	}

	return id, nil
}
