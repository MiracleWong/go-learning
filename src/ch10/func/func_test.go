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

// 可变参数的函数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParamas(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear Resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	// panic("err")
	// fmt.Println("End") // 未执行到的语句
}
