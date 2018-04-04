package palindrome

func IsA(subject string) bool {
	characters := []rune(subject)

	left, right := 0, len(characters) - 1

	for left < right {
		if characters[left] != characters[right] {
			return false
		}

		left++
		right--
	}

	return true
}
