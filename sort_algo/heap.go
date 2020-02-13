package sort_algo

// 堆排序
// 不稳定
// nlogn

func HeapSort(arr []int) []int {
	buildHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjust(arr[:i], 0)
	}
	return arr
}

func adjust(arr []int, parent int) {
	for {
		c := parent*2 + 1
		if c+1 < len(arr) && arr[c] < arr[c+1] {
			c++
		}
		// 没有子节点 || 满足最大堆
		if c >= len(arr) || arr[parent] > arr[c] {
			break
		}
		arr[parent], arr[c] = arr[c], arr[parent]
		parent = c
	}
}

func buildHeap(arr []int) {
	// 初始化
	for n := len(arr)/2 - 1; n >= 0; n-- {
		adjust(arr, n)
	}
}
