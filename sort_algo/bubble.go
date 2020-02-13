package sort_algo

// 冒泡排序
// 稳定
// n2

func BubbleSort(arr []int) []int {
	l := len(arr)
	for topN := 0; topN < l; topN++ {
		for next := 0; next < l-topN-1; next++ {
			if arr[next] > arr[next+1] {
				arr[next], arr[next+1] = arr[next+1], arr[next]
			}
		}
	}
	return arr
}
