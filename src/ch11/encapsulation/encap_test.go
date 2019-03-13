package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

// func (e Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID: %s-Name: %s-Age: %d", e.ID, e.Name, e.Age)
// }

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s / Name: %s / Age: %d", e.ID, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Logf("e is %T", e)
	t.Logf("e1 is %T", e1)
	// t.Logf("e1 is %T", &e1)
	t.Logf("e2 is %T", e2)
}

func TestEmployeeOprations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
