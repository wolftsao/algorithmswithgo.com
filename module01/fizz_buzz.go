package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			out[i-1] = "Fizz Buzz"
		case i%3 == 0:
			out[i-1] = "Fizz"
		case i%5 == 0:
			out[i-1] = "Buzz"
		default:
			out[i-1] = strconv.Itoa(i)
		}
	}

	fmt.Println(strings.Join(out, ", "))
}
