package main

import (
	"fmt"
)

type linearDict struct {
	data                []int
	actualSize, maxSize int
}

const N, D = 8, 10

func NewLinearDict() *linearDict {

	var obj = linearDict{
		data:       make([]int, defaultMaxSize),
		actualSize: 0,
		maxSize:    defaultMaxSize,
	}

	for i := range obj.data {
		obj.data[i] = -1
	}

	return &obj
}

func (ld *linearDict) isEmpty(i int) bool {
	return ld.data[i] == -1
}

func (ld *linearDict) isDeleted(i int) bool {
	return ld.data[i] == -2
}

func (ld *linearDict) scanFor(elem int) (int, int) {
	hashed := hash(elem, ld.maxSize)
	d := -1

	count := 1
	i := hashed

	for ; !ld.isEmpty(i); count++ {

		if count++; ld.isDeleted(i) {
			if d == -1 {
				d = i
			}

		} else if ld.data[i] == elem {
			return i, count
		}

		i = (i + 1) % ld.maxSize

		if i == hashed {
			return d, count
		}
	}

	if d != -1 {
		return d, count
	}

	return i, count
}

func (ld *linearDict) Find(toFind int) (int, bool, int) {
	i, count := ld.scanFor(toFind)

	count++
	if i == -1 || ld.isEmpty(i) || ld.isDeleted(i) {
		return -1, false, count
	}

	return ld.data[i], true, count
}

func (ld *linearDict) Insert(toInsert int) bool {
	i, _ := ld.scanFor(toInsert)

	if ld.isEmpty(i) || ld.isDeleted(i) {

		ld.data[i] = toInsert
		ld.actualSize++

		if ld.actualSize*D > ld.maxSize*N {
			ld.maxSize *= 2

			ld.data = ld.newData(ld.maxSize)

		}

		return true

	} else {
		ld.data[i] = toInsert

		return false
	}

}

func (ld *linearDict) Remove(toDelete int) {
	i, _ := ld.scanFor(toDelete)

	if !ld.isEmpty(i) && !ld.isDeleted(i) {

		ld.data[i] = -2
		ld.actualSize--

		if ld.maxSize > 1 && ld.actualSize*4*D < ld.maxSize*N {
			ld.maxSize /= 2
			ld.data = ld.newData(ld.maxSize)
		}
	}
}

func (ld *linearDict) newData(newSize int) []int {
	r := make([]int, newSize)

	for i := range r {
		r[i] = -1
	}

	for src_i, elem := range ld.data {

		if !ld.isEmpty(src_i) && !ld.isDeleted(src_i) {

			k := elem
			i := hash(k, newSize)

			for r[i] != -1 {

				i = (i + 1) % newSize

			}

			r[i] = ld.data[src_i]
		}

	}

	return r
}

/*

func (ld *linearDict) Insert(elem int) bool {

	hashed := hash(elem, ld.maxSize)

	for i := hashed; i < ld.maxSize; i++ {

		if ld.data[i] == elem {
			return false
		}

		if ld.data[i] == -1 {
			ld.data[i] = elem
			ld.actualSize++
			return true
		}
	}

	for i := 0; i < hashed; i++ {
		if ld.data[i] == elem {
			return false
		}

		if ld.data[i] == -1 {
			ld.data[i] = elem
			ld.actualSize++
			return true
		}
	}

	return false
}

func (ld *linearDict) Find(elem int) bool {

	hashed := hash(elem, ld.maxSize)

	for i := hashed; i < len(ld.data); i++ {

		if ld.data[i] == elem {
			return true
		}

	}

	for i := 0; i < hashed; i++ {

		if ld.data[i] == elem {
			return true
		}
	}

	return false
}

func (ld *linearDict) Remove(elem int) {

	hashed := hash(elem, ld.maxSize)

	for i := hashed; i < len(ld.data); i++ {

		if ld.data[i] == elem {
			ld.data[i] = -1
			return
		}

	}

	for i := 0; i < hashed; i++ {

		if ld.data[i] == elem {
			ld.data[i] = -1
			return
		}
	}
}*/

func (ld *linearDict) Show() {
	fmt.Println(ld.data)
}
