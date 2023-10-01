package main

func hash(toHash, numbOfIndex int) int {
	return toHash % numbOfIndex
}

const defaultMaxSize = 5

type dict interface {
	Find(toFind int) (int, bool, int)
	Insert(toAdd int) bool
	Remove(toRemove int)
	// Rehash()
}
