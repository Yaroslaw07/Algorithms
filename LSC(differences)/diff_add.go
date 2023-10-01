package main

type changeType int64

const (
	Added changeType = iota
	Removed
	Equal
	None
)

type changedCell struct {
	changeType
	value byte
}

func Addition(r byte) changedCell {
	return changedCell{Added, r}
}

func Removal(r byte) changedCell {
	return changedCell{Removed, r}
}

func Unchanged(r byte) changedCell {
	return changedCell{Equal, r}
}

func reverseWord(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func reverseSlice(slice []changedCell) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
