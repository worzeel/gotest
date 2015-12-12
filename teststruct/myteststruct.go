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
