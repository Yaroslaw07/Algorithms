package main

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return i / 2
}

type Heap struct {
	data []elem
}

func NewHeap() *Heap {
	return &Heap{[]elem{}}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) compare(i, j int) bool {
	return h.data[i].weight > h.data[j].weight
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Peek() (elem, bool) {
	if h.Size() <= 0 {
		return elem{}, false
	}

	return h.data[0], true
}

func (h *Heap) Push(e elem) {
	h.data = append(h.data, e)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap) heapifyUp(i int) {
	for h.compare(parent(i), i) {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func (h *Heap) Pop() (elem, bool) {

	if h.Size() < 0 {
		return elem{}, false
	}

	val := h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.heapifyDown(0)

	return val, true

}

func (h *Heap) heapifyDown(i int) {
	l, r := leftChild(i), rightChild(i)
	minimal := i

	if l < len(h.data) && h.compare(i, l) {
		minimal = l
	}

	if r < len(h.data) && h.compare(minimal, r) {
		minimal = r
	}

	if minimal != i {
		h.swap(i, minimal)
		h.heapifyDown(minimal)
	}
}
