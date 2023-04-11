package main

import (
	"fmt"
	"strconv"
)

type Animal struct {
	Name string
	Age  int
}

type canCall interface {
	call(string) string
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func (d *Dog) call(s string) string {
	return d.Name + strconv.Itoa(d.Age) + "is calling" + "from" + s
}

func (c *Cat) call(s string) string {
	return c.Name + strconv.Itoa(c.Age) + "is calling" + "from" + s
}

func doCall(o canCall, s string) {
	fmt.Println(o.call(s))
}

func main() {
	cat := new(Cat)
	cat.Name = "CAT001"
	cat.Age = 1222
	dog := Dog{Animal{Name: "dog001", Age: 22}}
	doCall(cat, "zzz")
	doCall(&dog, "zzz")
}
