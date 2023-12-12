package main


type MinHeap struct {
    elements []*HeapNode
}

type HeapNode struct {
    vertex string
    dist   float64
}

func NewMinHeap() *MinHeap {
    return &MinHeap{elements: make([]*HeapNode, 0)}
}

func (h *MinHeap) Push(node *HeapNode) {
    h.elements = append(h.elements, node)
    h.heapifyUp(len(h.elements) - 1)
}

func (h *MinHeap) Pop() *HeapNode {
    if len(h.elements) == 0 {
        return nil
    }

    n := len(h.elements) - 1
    h.swap(0, n)
    min := h.elements[n]
    h.elements = h.elements[:n]
    h.heapifyDown(0)

    return min
}

func (h *MinHeap) IsEmpty() bool {
    return len(h.elements) == 0
}

func (h *MinHeap) swap(i, j int) {
    h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *MinHeap) heapifyUp(index int) {
    for h.elements[parent(index)].dist > h.elements[index].dist {
        h.swap(parent(index), index)
        index = parent(index)
    }
}

func (h *MinHeap) heapifyDown(index int) {
    lastIndex := len(h.elements) - 1
    l, r := left(index), right(index)

    childToCompare := 0
    for l <= lastIndex {
        if l == lastIndex {
            childToCompare = l
        } else if h.elements[l].dist < h.elements[r].dist {
            childToCompare = l
        } else {
            childToCompare = r
        }

        if h.elements[index].dist > h.elements[childToCompare].dist {
            h.swap(index, childToCompare)
            index = childToCompare
            l, r = left(index), right(index)
        } else {
            return
        }
    }
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int  { return 2*i + 1 }
func right(i int) int { return 2*i + 2 }
