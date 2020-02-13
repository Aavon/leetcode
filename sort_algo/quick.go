package sort_algo

// 快速排序
// 不稳定
// nlogn

func QuickSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	left := 0
	right := l - 1
	split := arr[left] // 选left,初始化时只需要考虑right
	lr := 0            // 左右查找控制
	for left < right {
		if lr == 0 {
			if arr[right] >= split {
				right--
			} else {
				arr[left], arr[right] = arr[right], arr[left]
				lr = 1
			}
		} else if lr == 1 {
			if arr[left] <= split {
				left++
			} else {
				arr[left], arr[right] = arr[right], arr[left]
				lr = 0
			}
		}
	}
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}
