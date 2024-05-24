package heapDS

import (
	"context"
)

type MaxHeapI interface {
	BuildMaxHeap(ctx *context.Context, val int)
	ExtractMax(ctx *context.Context) int
	MaxHeapifyUp(ctx *context.Context, idx int)
	MaxHeapifyDown(ctx *context.Context, idx int)
	Swap(ctx *context.Context, i1, i2 int)
}
