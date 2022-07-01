package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GerFibonacciSeries(10))
	t.Log(series.Square(20))
}
