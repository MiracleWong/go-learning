package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func retureMultiValues() (int, int) {
	return rand.Intn(1), rand.Intn(20)
}
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func SlowFunc(op int) int {
	time.Sleep(time.Second * 10)
	return op
}
func TestFn(t *testing.T) {
	a, b := retureMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(SlowFunc)
	t.Log(tsSF(10))
}
