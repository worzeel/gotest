package teststruct

import "testing"

func TestAdd(t *testing.T) {
	myTest := MyTestStruct{1, 2}

	if myTest.Add() != 3 {
		t.Error("not three!")
	} else {
		t.Log("add is 3!")
	}
}
