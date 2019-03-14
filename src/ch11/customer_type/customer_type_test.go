package customer

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
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
	tsSF := timeSpent(SlowFunc)
	t.Log(tsSF(10))
	// t.Log("Hello World")
}
