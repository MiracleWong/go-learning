package main

import "fmt"

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  uint
	addr address
}

// Factory Func can auto generate:
// Generate(Alt + Insert) --> Constructor
func newPerson(name string) *person {
	return &person{name: name}
}

type Stringer interface {
	// 通过它的 String() 方法获取一个字符串

	String() string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d ", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the ddar is %s %s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter是Reader和Writer的组合

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	p := person{
		name: "MiracleWong",
		age:  30,
		addr: address{
			province: "Zhejiang",
			city:     "Hangzhou",
		},
	}

	fmt.Println(p.addr.province, p.addr.city)

	printString(p)
	printString(&p)
}
