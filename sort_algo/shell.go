package sort_algo

// 希尔排序
// 不稳定
// nlogn

func ShellSort(arr []int) []int {
	l := len(arr)
	for step := l / 2; step > 0; step = step / 2 {
		for groupStart := 0; groupStart < step; groupStart++ {
			for bottomN := groupStart; bottomN < l-step; bottomN += step {
				for next := bottomN + step; next > groupStart; next -= step {
					if arr[next-step] > arr[next] {
						arr[next-step], arr[next] = arr[next], arr[next-step]
					} else {
						break
					}
				}
			}
		}
	}
	return arr
}
