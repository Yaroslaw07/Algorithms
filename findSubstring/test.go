package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getLetter(size int) string {
	return fmt.Sprintf("%c", rand.Intn(size)+65)
}

func getWord(length int, alphabetSize int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteString(getLetter(alphabetSize))
	}

	return sb.String()
}

func addLetter(word string, alphabetSize int) string {
	return word + getLetter(alphabetSize)
}

func WriteResults(nameOfFile string, results [][4]int) {

	f, err := os.Create("./results/" + nameOfFile)

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range results {
		stringRow := make([]string, len(record))
		for i, num := range record {
			stringRow[i] = strconv.Itoa(num)
		}
		// Write the row to the CSV file
		w.Write(stringRow)
	}
}

func testText() {
	pattern := getWord(5, 26)
	text := ""

	results := [][4]int{}
	for i := 0; i < 100; i++ {

		for j := 0; j < 100; j++ {
			text = addLetter(text, 26)
		}

		results = append(results, [4]int{})
		results[i][0] = (i + 1) * 100
		results[i][1], _ = FindSimple(text, pattern)
		results[i][2], _ = FindSunday(text, pattern)
		results[i][3], _ = FindRabinKarp(text, pattern)
	}

	WriteResults("test-text.csv", results)
}

func testPattern() {
	pattern := ""
	text := getWord(10000, 26)

	results := [][4]int{}
	for i := 0; i < 40; i++ {

		pattern += getLetter(26)

		results = append(results, [4]int{})
		results[i][0] = i + 1
		results[i][1], _ = FindSimple(text, pattern)
		results[i][2], _ = FindSunday(text, pattern)
		results[i][3], _ = FindRabinKarp(text, pattern)
	}

	WriteResults("test-pattern.csv", results)
}

func testAlphabet() {

	results := [][4]int{}

	for i := 0; i < 20; i++ {

		text := getWord(10000, i+1)
		pattern := getWord(10, i+1)

		results = append(results, [4]int{})
		results[i][0] = (i + 1)
		results[i][1], _ = FindSimple(text, pattern)
		results[i][2], _ = FindSunday(text, pattern)
		results[i][3], _ = FindRabinKarp(text, pattern)

	}

	WriteResults("test-alphabet.csv", results)
}
