package s5

import (
	"fmt"
	"sort"
	"testing"
)

// 123542611, 交换数字顺序，求最小的整数

func solution(val []int) []int {
	exchangeA, exchangeB := -1, -1

	for i := len(val) - 1; i > 0; i-- {
		if exchangeB == -1 || val[i] >= val[exchangeB] {
			exchangeB = i
		} else if val[i] < val[exchangeB] {
			exchangeA = i
			break
		}
	}

	if exchangeA == -1 {
		return val
	}

	// nestest
	sort.Ints(val[exchangeA+1:])

	for i := exchangeA + 1; i < len(val); i++ {
		if val[exchangeA] < val[i] {
			exchangeB = i
			break
		}
	}

	val[exchangeA], val[exchangeB] = val[exchangeB], val[exchangeA]
	return val
}

func Test_a(t *testing.T) {
	val := []int{1, 2, 3, 5, 4, 2, 6, 1, 1}
	fmt.Println(val)
	fmt.Println(solution(val))
}
