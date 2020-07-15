package State

import "testing"

func TestSate(t *testing.T)  {
	machine := NewMachine()
	machine.Off()  //一开始OFF的状态，调用OFF结构体中的on和off方法
	machine.On()
	machine.On()  //上面已经是on状态，那么便调用ON结构体中的on和off
	machine.Off()
}
