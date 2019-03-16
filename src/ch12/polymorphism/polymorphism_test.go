package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Pringln(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "system.out.Pringln(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}
func TestPolymorphism(t *testing.T) {
	goPro := new(GoProgrammer)
	javaPro := new(JavaProgrammer)
	writeFirstProgram(goPro)
	writeFirstProgram(javaPro)
}
