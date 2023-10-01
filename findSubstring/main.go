package main

import (
	"os"
)

func main() {

	os.Mkdir("results", 0755)
	testText()
	testPattern()
	testAlphabet()

	/*fmt.Println(FindSunday("ava", "ava"))
	fmt.Println(FindSunday("The quick brown fox jumps over the lazy dog.", "brown"))
	fmt.Println(FindSunday("The cat in the hat sat on the mat.", "cat"))
	fmt.Println(FindSunday("Peter Piper picked a peck of pickled peppers.", "peck"))
	fmt.Println(FindSunday("How much wood would a woodchuck chuck, if a woodchuck could chuck wood?", "wood"))*/
}
