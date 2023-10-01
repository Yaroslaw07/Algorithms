package main

func matchesAt(source string, pos int, toFind string) (int,bool) {

	count := 0

	for i := 0; i < len(toFind);i++ {

		if count++; source[pos + i] != toFind[i] {
			return count,false
		}
	}

	return count,true
}

func FindSimple(source, toFind string) (int, []int) {

	res := []int{}
	count := 0

	for i := 0; i < len(source)-len(toFind)+1; i++ {

		addCount, ok := matchesAt(source, i, toFind)
		count += addCount
		if ok {
			res = append(res, i)
		}
	}

	return count, res
}

func FindSunday(source, toFind string) (int, []int) {

	tab := map[rune]int{}

	for i, elem := range toFind {

		tab[elem] = i
	}

	res := []int{}
	count := 0

	for i := 0; i < len(source)-len(toFind)+1; {

		addCount, ok := matchesAt(source, i, toFind)
		count += addCount
		if ok{
			res = append(res, i)
		}

		i += len(toFind)

		if i < len(source) {

			if step, ok := tab[rune(source[i])]; ok {
				i -= step
			} else {
				i += 1
			}

		}

	}

	return count, res
}

func FindRabinKarp(source, toFind string) (int, []int) {

	const q, d = 21, 256

	height := 1

	for i := 0; i < len(toFind)-1; i++ {
		height = (height * d) % q
	}

	patternHash, sourceHash := 0, 0
	res := []int{}
	count := 0

	for i := 0; i < len(toFind) && i < len(source); i++ {
		patternHash = (d*patternHash + int(toFind[i])) % q
		sourceHash = (d*sourceHash + int(source[i])) % q
	}

	for i := 0; i < len(source)-len(toFind)+1; i++ {

		if patternHash == sourceHash {

			addCount, ok := matchesAt(source, i, toFind)
			count += addCount	

			if ok{
				res = append(res, i)
			}
		}

		if i < len(source)-len(toFind) {
			sourceHash = (d*(sourceHash-int(source[i])*height) + int(source[i+len(toFind)])) % q

			if sourceHash < 0 {
				sourceHash += q
			}

		}
	}

	return count, res
}
