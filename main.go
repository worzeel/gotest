// main.go
package main

import (
	"fmt"
)

type MyTest struct {
	name string
	age int
}

func (self MyTest) random() MyTest {
	self.name = "Random"
	//self.age = 123
	
	return self
}

func (self MyTest) getThings() (string, int) {
	return "from GetThings", 11
}

func main() {
	test := MyTest{}
	
	test.age = 111
	test.name = "test"
	
	test2 := test.random()
	
	getThingsStr, getThingsInt := test.getThings()
	
	fmt.Println("testing ", test.name, " ", test.age, " ", test2.name, " ", test2.age)
	fmt.Println(getThingsStr, " - ", getThingsInt)
}
