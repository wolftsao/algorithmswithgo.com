package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	if base > 16 {
		return 0
	}

	res := 0
	for i, v := range value {
		d, err := strconv.ParseInt(string(v), 16, 64)
		if err != nil {
			return 0
		}

		res += int(math.Pow(float64(base), float64(len(value)-1-i))) * int(d)
	}

	return res
}
