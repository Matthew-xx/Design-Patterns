package mediator

import "testing"

func TestNewMediator(t *testing.T) {
	mediator := NewMediator()  //新建中介，其下右ted和bill
	mediator.Ted.Talk()  //Ted说，然后中介就去找bill(如果是bill说那么便去找ted回应）Communicate相当于一个调度
}
