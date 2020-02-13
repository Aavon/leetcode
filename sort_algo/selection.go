package sort_algo

// 选择排序
// 不稳定
// n2

func SelectionSort(arr []int) []int {
	l := len(arr)
	for bottomN := 0; bottomN < l-1; bottomN++ {
		for next := bottomN + 1; next < l; next++ {
			if arr[bottomN] > arr[next] {
				arr[bottomN], arr[next] = arr[next], arr[bottomN]
			}
		}
	}
	return arr
}
