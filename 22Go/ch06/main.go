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

func main() {
	p := person{
		name: "MiracleWong",
		age:  30,
		addr: address{
			province: "Zhejiang",
			city:     "Hangzhou",
		},
	}

	fmt.Println(p.addr.city)
}
