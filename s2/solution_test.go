package s2

import (
	"net"
	"testing"
)

func Test_a(t *testing.T) {

}

func GetMaxRepeatNum(a []int) int {
	stat := make(map[int]int)
	for _, v := range a {
		stat[v]++
	}

	maxRepeatVal := 0
	maxRepeatCnt := 0
	for k, v := range stat {
		if v > maxRepeatCnt {
			maxRepeatCnt = v
			maxRepeatVal = k
		}
	}
	return maxRepeatVal
}

func GetMaxRepeatNum2(a []int) int {
	conn, _ := net.Dial("", "")
	conn.Read()
}
