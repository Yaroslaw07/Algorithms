package main

import "fmt"

type listDict struct {
	data                [][]int
	actualSize, maxSize int
}

const maxPerList = 20

func NewListDict() *listDict {

	var obj listDict = listDict{
		data:       make([][]int, defaultMaxSize),
		actualSize: 0,
		maxSize:    defaultMaxSize}

	return &obj
}

func (ld *listDict) Insert(toInsert int) bool {

	h, i, _ := ld.find_index(toInsert)

	if i == -1 {
		ld.data[h] = append(ld.data[h], toInsert)
		ld.actualSize++

		if ld.actualSize > ld.maxSize*maxPerList {
			ld.maxSize *= 2
			ld.data = ld.newData(ld.maxSize)

		}

		return true

	} else {
		ld.data[h][i] = toInsert
		return false
	}
}

func (ld *listDict) find_index(toFind int) (int, int, int) {

	h := hash(toFind, ld.maxSize)
	count := 0

	for i, elem := range ld.data[h] {
		if count++; toFind == elem {
			return h, i, count
		}
	}

	return h, -1, count
}

func (ld *listDict) Find(toFind int) (int, bool, int) {

	h, i, count := ld.find_index(toFind)

	if i == -1 {
		return -1, false, count
	}

	return ld.data[h][i], true, count
}

func (ld *listDict) Remove(toRemove int) {

	h, i, _ := ld.find_index(toRemove)

	if i != -1 {
		ld.data[h] = append(ld.data[h][:i], ld.data[h][i+1:]...)
		ld.actualSize--

		if ld.maxSize > 1 && 4*ld.actualSize < ld.maxSize*maxPerList {
			ld.maxSize /= 2
			ld.data = ld.newData(ld.maxSize)

		}
	}
}

func (ld *listDict) newData(newSize int) [][]int {

	r := make([][]int, newSize)

	for _, subArr := range ld.data {

		for _, elem := range subArr {

			h := hash(elem, newSize)

			r[h] = append(r[h], elem)
		}

	}

	return r
}

func (sd *listDict) Show() {

	for i, subArr := range sd.data {

		fmt.Print(i, ": ")

		for _, elem := range subArr {
			fmt.Print(elem, "| ")
		}

		fmt.Println()
	}

	fmt.Println("-----------------------------")
}
