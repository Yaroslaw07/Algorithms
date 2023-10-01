package main

import (
	"fmt"
)

func main() {
	gra := [][]node{
		0: {{1, 3}, {2, 4}},
		1: {{2, 3}, {3, 8}},
		2: {{4, 1}, {3, 7}},
		3: {{2, 4}, {0, -1}},
		4: {{1, 2}, {3, -1}},
	}

	FloydWarshall(gra)
	//dist[][] will be the output matrix that will finally
	//have the shortest distances between every pair of vertices
	for i, d := range dict {
		fmt.Printf("%d -- %4g\n", i, d)
	}

	fmt.Println("-------------------------------")

	for i := range dict {
		for j := range dict {
			fmt.Println(i, "->", j, ": ", findWay(i, j), " $$ ", dict[i][j])
		}
	}
}
