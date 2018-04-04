package substrings

type window struct {
	start, end int
}

func (w window) in(i int) bool {
	return w.start <= i && i <= w.end
}

func (w window) len() int {
	return w.end - w.start + 1
}

func LongestUnique(subject string) string {
	charactersByLastIndex := map[rune]int{}

	var maxWindow window
	var currentWindow window

	characters := []rune(subject)

	for i, c := range characters {
		if lastSeenAt, present := charactersByLastIndex[c]; !present {
			charactersByLastIndex[c] = i
			currentWindow.end++
		} else if currentWindow.in(lastSeenAt) {
			w := window{lastSeenAt+1, i}

			if currentWindow.len() > maxWindow.len() {
				maxWindow = currentWindow
			}

			currentWindow = w
		}
	}

	if currentWindow.len() > maxWindow.len() {
		maxWindow = currentWindow
	}

	return subject[maxWindow.start : maxWindow.end]
}
