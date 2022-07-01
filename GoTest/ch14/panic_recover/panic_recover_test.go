package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("recoverd panic", error)
		}
		fmt.Println("Finally!")
	}()
	fmt.Println("start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
