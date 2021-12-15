package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// runes := []rune(word)

	// for i := 0; i < len(runes)/2; i++ {
	// 	runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	// }

	// return string(runes)

	var res string

	for _, r := range word {
		res = string(r) + res
	}
	return res
}
