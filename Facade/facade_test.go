package Facade

import "testing"

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	f := CarFacade{}  //调用门面的结构体，再调用里面的部件
	f.CreateCompleteCar()
}
