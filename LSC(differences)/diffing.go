package main

func compute_lcs(text1, text2 string) [][]int {

	n, m := len(text1), len(text2)

	lcs := [][]int{}

	for i := 0; i <= n; i++ {

		lcs = append(lcs, []int{})

		for j := 0; j <= m; j++ {

			if i == 0 || j == 0 {
				lcs[i] = append(lcs[i], 0)
				continue
			}

			if text1[i-1] == text2[j-1] {
				lcs[i] = append(lcs[i], 1+lcs[i-1][j-1])
				continue
			}

			if lcs[i-1][j] > lcs[i][j-1] {
				lcs[i] = append(lcs[i], lcs[i-1][j])
			} else {
				lcs[i] = append(lcs[i], lcs[i][j-1])
			}
		}

	}

	return lcs

}

func find_lcs(text1, text2 string) string {
	result := ""

	lcs := compute_lcs(text1, text2)

	i, j := len(text1), len(text2)

	for i != 0 && j != 0 {

		if text1[i-1] == text2[j-1] {
			result = result + string(text1[i-1])
			i, j = i-1, j-1
			continue
		}

		if lcs[i-1][j] <= lcs[i][j-1] {
			j -= 1
			continue
		}

		i -= 1
	}

	return reverseWord(result)
}

func diff(text1, text2 string) []changedCell {
	lcs := compute_lcs(text1, text2)

	results := []changedCell{}

	i, j := len(text1), len(text2)

	for i != 0 || j != 0 {

		if i == 0 {
			results = append(results, Addition(text2[j-1]))
			j -= 1
			continue
		}

		if j == 0 {
			results = append(results, Removal(text1[i-1]))
			i -= 1
			continue
		}

		if text1[i-1] == text2[j-1] {
			results = append(results, Unchanged(text1[i-1]))
			i, j = i-1, j-1
			continue
		}

		if lcs[i-1][j] <= lcs[i][j-1] {
			results = append(results, Addition(text2[j-1]))
			j -= 1
			continue
		}

		results = append(results, Removal(text1[i-1]))
		i -= 1
	}

	reverseSlice(results)

	return results
}
