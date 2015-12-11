// Package teststruct - myteststruct.go
package teststruct

// MyTestStruct test
type MyTestStruct struct {
	X int
	Y int
}

// SetValues testing
func (m *MyTestStruct) SetValues(x, y int) {
	m.X = x
	m.Y = y
}

// Testing is a test
func (m MyTestStruct) Testing() int {
	return m.X + m.Y
}
