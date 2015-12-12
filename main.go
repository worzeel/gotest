// gotest
package main

import (
	"fmt"
	"gotest/interfaces"
	"gotest/teststruct"
)

func myTestFunc(x, y int) int {
	return x + y
}

func myTestFunc2(x, y int) (int, string) {
	return x + y, "test"
}

// this type of return should only be used for very small methods for readability
func myTestNamedFunc() (x, y int) {
	x = 3
	y = 5

	// naked return
	return
}

func myTestInterfaceUse(m interfaces.MyTest) {
	fmt.Println("Test interface:", m.MyOutput())
}

func main() {
	var aTest string
	aTest = "test"

	a, b := myTestFunc2(2, 3)
	c, d := myTestNamedFunc()

	fmt.Println("this is just a testing", myTestFunc(1, 2))
	fmt.Println("second:", a, b)
	fmt.Println("third:", c, d)
	fmt.Println("fourth:", aTest)

	sum := 0
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("fifth:", sum)

	switch test := 2; test {
	case 1:
		fmt.Println(1)

	case 2:
		fmt.Println(2)

	default:
		fmt.Println(test)
	}

	switch {
	case a == 5:
		fmt.Println("five")

	default:
		fmt.Println("something else")
	}

	myTest := teststruct.MyTestStruct{}
	myTest.SetValues(1, 2)
	fmt.Println("Add values:", myTest.Add())
	fmt.Println("Multiple values:", myTest.Multiply())

	fmt.Println("Struct test", teststruct.MyTestStruct{1, 2})

	i := 1
	p := &i
	j := i + *p

	fmt.Println("i:", i, "*p:", *p, "p:", p)
	fmt.Println("j:", j)

	myTestInterfaceUse(myTest)

	/* This will _not_ work, as MyTestStruct2 does not implement the MyOutput method!
	myTest2 := teststruct.MyTestStruct2{}
	myTestInterfaceUse(myTest2)
	*/
}
