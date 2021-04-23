package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func NewHuman() (h *Human) {
	h = new(Human)
	return
}

func NewHumanWithParams(firstname, lastname string, age int) (h *Human) {
	h = new(Human)
	h.FirstName = firstname
	h.LastName = lastname
	h.Age = age
	return
}

func main() {
	//definition v1
	// x := Human{FirstName: "Rafet"}
	// fmt.Println(x.FirstName)

	//definition v2
	// x := new(Human)
	// x.FirstName = "Rafet"
	// fmt.Println(x.FirstName)

	// definition v3 constructor kullanarak
	// x := NewHuman()
	// x.FirstName = "Rafet"
	// fmt.Println(x.FirstName)

	// definition v4 constructor parameter kullanarak
	x := NewHumanWithParams("rafet", "parlak", 25)
	y := x
	y.FirstName = "Serenay"
	var outData = x.FirstName + " " + x.LastName + " / " + strconv.Itoa(x.Age)
	var outData2 = y.FirstName + " " + y.LastName + " / " + strconv.Itoa(y.Age)
	fmt.Println(outData)
	fmt.Println(outData2)
}
