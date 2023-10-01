package main

import (
	"fmt"
	"math"
)

type node struct {
	to int
	wt float64
}

var dict [][]float64
var ways [][]int

func FloydWarshall(n [][]node) {
	dict = make([][]float64, len(n))
	ways = make([][]int, len(n))

	for i := range n {

		di := make([]float64, len(n))
		way := make([]int, len(n))

		for j := range di {
			di[j] = math.Inf(1)
			way[j] = -1
		}

		di[i] = 0
		dict[i], ways[i] = di, way
	}

	for u, graphs := range n {
		for _, v := range graphs {

			dict[u][v.to] = v.wt
			ways[u][v.to] = u
		}
	}

	for k := range dict {
		for i := range dict {
			for j := range dict {

				if d := dict[i][k] + dict[k][j]; d < dict[i][j] {
					dict[i][j] = d
					ways[i][j] = ways[k][j]
				}
			}
		}
	}
}

func findWay(i, j int) (path string) {

	if i == j {
		return fmt.Sprint(i) + " "
	}

	if ways[i][j] == -1 {
		return "NO PATH"
	}

	return findWay(i, ways[i][j]) + " " + fmt.Sprint(j)
}
