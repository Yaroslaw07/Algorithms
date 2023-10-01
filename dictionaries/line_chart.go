package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var testSets = map[string]dict{
	"Linear": NewLinearDict(),
	"List":   NewListDict(),
}

const numbOfTests = 10000

func init() {
	rand.Seed(time.Now().UnixNano())
}

func makeTests() {

	results := [][3]int{}
	linear := NewLinearDict()
	list := NewListDict()

	for i := 0; i < numbOfTests/10; i++ {

		for j := 0; j < 9; j++ {
			getData(linear)
			getData(list)
		}

		results = append(results, [3]int{})
		results[i][0] = (i + 1) * 10
		results[i][1] = getData(linear)
		results[i][2] = getData(list)
	}

	WriteResults("results.csv", results)
}

func WriteResults(nameOfFile string, results [][3]int) {

	f, err := os.Create(nameOfFile)

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

func getData(d dict) int {

	ok := false
	for !ok {
		ok = d.Insert(rand.Intn(10000000000))
	}

	_, _, count := d.Find(rand.Intn(10000000000))

	return count

}

/*func makeTests() {

	axis := []int{}

	for i := 1; i <= numbOfTests; i++ {
		axis = append(axis, i)
	}

	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Compare two maps by count finds",
		}),
	)

	line.SetXAxis(axis)

	for name, dict := range testSets {
		line.AddSeries(name, getData(dict))
	}

	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, _ := os.Create("res.html")
	_ = line.Render(f)

}

func getData(d dict) []opts.LineData {

	items := make([]opts.LineData, 0)

	for i := 0; i < numbOfTests; i++ {

		ok := false
		for !ok {
			ok = d.Insert(rand.Intn(10000))
		}

		_, _, count := d.Find(rand.Intn(10000))
		items = append(items, opts.LineData{Value: count})
	}

	return items
}
*/
