// Package teststruct - myteststruct.go
package teststruct

// MyTestStruct test
type MyTestStruct struct {
	X int
	Y int
}

// SetValues - Call a method to do stuff
func (m *MyTestStruct) SetValues(x, y int) {
	m.X = x
	m.Y = y
}

// Add - add the two values
func (m MyTestStruct) Add() int {
	return m.X + m.Y
}

// Multiply - multiply the two values
func (m MyTestStruct) Multiply() int {
	return m.X * m.Y
}

// MyOutput - test output to adhear to MyTest interface
func (m MyTestStruct) MyOutput() string {
	return "output from MyTestStruct"
}

// MyTestStruct2 test 2
type MyTestStruct2 struct {
	//
}
