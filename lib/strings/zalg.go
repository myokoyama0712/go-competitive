package strings

// ZAlgorithm calculates Z-values of a slice S.
// Z-values show a maximum length of prefix between S and S[i:] for each i.
// Time complexity: O(n)
func ZAlgorithm(S []rune) (Z []int) {
	_min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	_max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(S)
	Z = make([]int, n)

	// Scanned subarray that has the same prefix in the most right.
	from, last := -1, -1
	for i := 1; i < n; i++ {
		same := &Z[i]

		// Z-Algorithm trick
		if from != -1 {
			*same = _min(Z[i-from], last-i)
			*same = _max(0, *same)
		}

		// scan naively
		for i+(*same) < n && S[*same] == S[i+(*same)] {
			(*same)++
		}

		// update the most right scanned subarray
		if last < i+(*same) {
			last = i + (*same)
			from = i
		}
	}

	Z[0] = n
	return Z
}
