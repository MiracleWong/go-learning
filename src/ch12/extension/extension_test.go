package panic_recover

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("   ", host)
}

type Dog struct {
	Pet
}

// type Dog struct {
// 	p *Pet
// }

func (d *Dog) Speak() {
	fmt.Print("汪汪汪")
}

//
// func (d *Dog) SpeakTo(host string) {
// 	d.Speak()
// 	fmt.Println("", host)
// }

func TestExtension(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Hello ")
}
