package palindrome

import (
	"strings"
)

func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}

		return strings.Split(s, "")[0]
	}

	input := strings.Split(s, "")
	main := make([]string, (len(input) * 2))
	lengths := make([]int, len(main))

	// we fill the main array with sentinel characters that represent the spaces between the numbers
	// since we are handling both even and odd numbered arrays
	for i := range main {
		if i == 0 {
			main[0] = input[0]
			continue
		}

		if i%2 != 0 {
			main[i] = "$"
			continue
		}

		main[i] = input[i/2]
	}

	// now with this setup, we can begin to fill our lengths array
	//   previous := make([]int, 2)

	for i := range main {
		lengths[i] = 1

		if i == 0 || i+1 == len(main) {
			continue
		}

		if main[i-1] != main[i+1] {
			continue
		}

		if main[i-1] == main[i+1] {
			find(main, lengths, i, i-1, i+1)

			newCenter(main, lengths, i)
		}

	}

	highPosition := 0
	for i := range lengths {
		if lengths[i] > lengths[highPosition] {
			highPosition = i
		}
	}

	lowBound := highPosition - ((lengths[highPosition] - 1) / 2)
	highBound := highPosition + ((lengths[highPosition] - 1) / 2)

	return strings.Replace(strings.Join(main[lowBound:highBound+1], ""), "$", "", -1)
}

func find(main []string, lengths []int, iter, back, front int) {
	if back < 0 || front > len(main)-1 {
		return
	}

	if main[back] != main[front] {
		return
	}

	lengths[iter] = lengths[iter] + 2
	find(main, lengths, iter, back-1, front+1)
}

func newCenter(main []string, lengths []int, iter int) int {
	// current palindrome bounds
	lowBound := iter - ((lengths[iter] - 1) / 2)
	highBound := iter + ((lengths[iter] - 1) / 2)

	/*
		We infer the next length in line from the previous, mirrored palindrome
	*/

	// X Contained in current palindrome?

	// X Palindrome extends to end of input?

	// Y Palindrome extends to right edge AND mirror palindrome is a prefix

	// X Expands to right edge, mirror palindrome expands past left, invalid

	return 0
}
