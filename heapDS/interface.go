package heapDS

import (
	"context"
)

type MaxHeapI interface {
	InsertMaxHeap(ctx *context.Context, val int) (MaxHeap, error)
	ExtractMax(ctx *context.Context) (int, error)
	MaxHeapifyUp(ctx *context.Context, idx int)
	MaxHeapifyDown(ctx *context.Context, idx int)
	Swap(ctx *context.Context, i1, i2 int)
}
