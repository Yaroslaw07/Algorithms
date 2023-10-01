package main

import "fmt"

func prim(nodes [][]node, start int) (res []elem, allCost int) {

	res, allCost = []elem{}, 0

	if start < 0 && len(nodes) >= start {
		return res, allCost
	}

	visited := make(map[int]bool)
	distances := NewHeap()

	// get all from start
	visited[start] = true

	for _, gr := range nodes[start] {
		distances.Push(elem{start, gr})
	}

	// main cycle
	for len(visited) < len(nodes) && distances.Size() != 0 {

		val, _ := distances.Pop()

		if visited[val.to] {
			continue
		}

		res = append(res, val)
		visited[val.to] = true
		allCost += val.weight

		for _, gr := range nodes[val.to] {

			if visited[gr.to] {
				continue
			}

			distances.Push(elem{val.to, gr})
		}
	}

	return res, allCost
}

func main() {
	gra := [][]node{
		0: {{1, 3}, {2, 4}},
		1: {{2, 3}, {3, 8}, {0, 1}},
		2: {{4, 1}, {3, 7}},
		3: {{2, 4}, {0, 5}},
		4: {{1, 2}, {3, 2}},
	}

	res, allCost := prim(gra, 2)

	for _, e := range res {
		fmt.Println(e.from, "--", e.weight, "->", e.to)
	}

	fmt.Println(allCost)
}
