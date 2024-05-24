package heapDS

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	MaxHeapList []int `json:"maxHeapList"`
}

// Insert adds an element to the heap
func (h *MaxHeap) BuildMaxHeap(ctx *gin.Context, list []int) {

	for _, val := range list {
		h.MaxHeapList = append(h.MaxHeapList, val)
		h.MaxHeapifyUp(ctx, len(h.MaxHeapList)-1)
	}
}

// Extracts returns the largets	key, and removes it from the heap
func (h *MaxHeap) ExtractMax(ctx *gin.Context, idx int) int {
	extracted := h.MaxHeapList[0]
	l := len(h.MaxHeapList) - 1

	// when the array is empty
	if len(h.MaxHeapList) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// take out the last element and put it in the root
	h.MaxHeapList[0] = h.MaxHeapList[l]
	h.MaxHeapList = h.MaxHeapList[:l]

	h.MaxHeapifyDown(ctx, 0)
	return extracted
}

// MaxHeapify will heapify from bottom top
func (h *MaxHeap) MaxHeapifyUp(ctx *gin.Context, idx int) {
	for h.MaxHeapList[parent(idx)] < h.MaxHeapList[idx] {
		h.Swap(ctx, parent(idx), idx)
		idx = parent(idx)
	}
}

func (h *MaxHeap) Swap(ctx *gin.Context, i1, i2 int) {
	h.MaxHeapList[i1], h.MaxHeapList[i2] = h.MaxHeapList[i2], h.MaxHeapList[i1]
}

// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) MaxHeapifyDown(ctx *gin.Context, idx int) {
	lastIndex := len(h.MaxHeapList) - 1
	l, r := left(idx), right(idx)
	childToCompare := 0

	// loop when index has atleast one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.MaxHeapList[l] > h.MaxHeapList[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.MaxHeapList[idx] < h.MaxHeapList[childToCompare] {
			h.Swap(ctx, idx, childToCompare)
			idx = childToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}
