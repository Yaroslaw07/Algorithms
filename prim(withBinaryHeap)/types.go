package main

type node struct {
	to     int
	weight int
}

type elem struct {
	from int
	node
}
