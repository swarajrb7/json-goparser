package main

func CompareRuneSlice(rune1 []rune, rune2 []rune, n int) bool {
	if n > len(rune1) || n > len(rune2) {
		return false
	}
	for i := 0; i < n; i++ {
		if rune1[i] != rune2[i] {
			return false
		}
	}
	return true
}