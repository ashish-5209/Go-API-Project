package heapDS

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}
