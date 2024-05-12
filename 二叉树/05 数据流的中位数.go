package main

import (
	"container/heap"
)

/*
	如何得到一个数据流中的中位数？

	如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
	如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

	例如，
	[2,3,4]的中位数是3
	[2,3]的中位数是(2+3)/2 = 2.5

	设计一个支持以下两种操作的数据结构：

	void addNum(int num) - 添加一个整数到数据结构中。
	double findMedian() - 返回目前所有元素的中位数。
*/

/*
	这道题目得先了解以下几个基础概念：

	1、中位数指的是排序数组的中间元素值，如果是奇数，那么直接就是中间的数值；如果是偶数，那么就是中间两个数的平均值。
	2、数据流指的是数据的长度是动态变化的，就像流水一样，在不断的新增数据过来。

	这意味着，数据流的中位数在不断的变化，不仅值在变化，求解方式也是在动态变化。

	并且，我们需要不断的将数据流中的全部数字进行排序，我们这里利用堆来实现。

	设置两个堆，一个是大顶堆maxHeap，一个是小顶堆minHeap。

	大顶堆maxHeap来存储数据流中较小一半的值
	小顶堆minHeap来存储数据流中较大一半的值

	由于大顶堆的堆顶为它的存储区间的最大值，小顶堆的堆顶为它的存储区间的最小值，那么如果用这两个堆来存储数据流的所有数据，我们可以组成一个递增有序的数组。

	1、大顶堆从左到右递增
	2、小顶堆从左到右递增

	在动态存储数据流的数据过程中，中位数也就是这两种情况：

	1、数据流为奇数时，保证两个堆的长度相差1，小顶堆的堆顶就是中位数。
	2、数据流为偶数时，保证两个堆的长度相等，两个堆的堆顶相加除二就是中位数。

	那一个新来的数据应该添加到哪个堆呢？

	1、当两堆长度相等，即长度都为n时，新数据加入到小顶堆中，使得小顶堆的长度为n+1，那么这两个堆的长度相差1，小顶堆的堆顶就是中位数。
	2、当两堆长度不相等，即小顶堆长度为n，大顶堆长度为n-1，新数据加入到大顶堆中，使得大顶堆的长度为n ，那么两个堆的长度就相等，两个堆的堆顶相加除二就是中位数。

	但是，这样做会有一个问题。新来的数据按上述方式插入后从大顶堆到小顶堆可能不是一个递增有序数组了。

	为了让每次新增一个数据到两个堆之后，都能使得从大顶堆到小顶堆是一个递增有序的数组，我们可以采取如下的操作。

	1、当两堆长度相等，即长度都为n时，新数据先加入到大顶堆中，然后再把大顶堆的堆顶元素取出，放入到小顶堆中。
	2、当两堆长度不相等，即小顶堆长度为n，大顶堆长度为n-1，新数据先加入到小顶堆中，然后再把小顶堆的堆顶元素取出，放入到大顶堆中。
*/

type HeapType int32

const (
	MinHeap HeapType = iota
	MaxHeap
)

type Heap struct {
	arr []int
	typ HeapType
}

func (h Heap) Len() int      { return len(h.arr) }
func (h Heap) Swap(i, j int) { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }
func (h Heap) Less(i, j int) bool {
	if h.typ == MinHeap {
		return h.arr[i] < h.arr[j]
	} else {
		return h.arr[i] > h.arr[j]
	}
}

func (h Heap) Top() int {
	return h.arr[0]
}

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	(*h).arr = append((*h).arr, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old.arr)
	x := old.arr[n-1]
	(*h).arr = old.arr[:n-1]
	return x
}

var (
	minHeap = &Heap{typ: MinHeap}
	maxHeap = &Heap{typ: MaxHeap}
)

func addNum(num int) {
	if maxHeap.Len() != minHeap.Len() {
		heap.Push(minHeap, num)
		heap.Push(maxHeap, heap.Pop(minHeap))
	} else {
		heap.Push(maxHeap, num)
		heap.Push(minHeap, heap.Pop(maxHeap))
	}
}

func findMedian() float64 {
	if maxHeap.Len() != minHeap.Len() {
		return float64(minHeap.Top())
	} else {
		return (float64(maxHeap.Top()) + float64(minHeap.Top())) / 2
	}
}

/*
	测试代码：
	func main() {
		heap.Init(minHeap)
		heap.Init(maxHeap)

		addNum(5)
		addNum(1)
		addNum(6)
		addNum(3)
		addNum(4)
		addNum(2)

		fmt.Println(findMedian())

		for maxHeap.Len() > 0 {
			fmt.Printf("%d ", heap.Pop(maxHeap))
		}
		for minHeap.Len() > 0 {
			fmt.Printf("%d ", heap.Pop(minHeap))
		}
	}
*/
