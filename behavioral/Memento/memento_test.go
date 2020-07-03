package Memento

import (
	"fmt"
	"testing"
)

func TestNumber_ReinstateMemento(t *testing.T) {
	n := NewNumber(7)
	n.Double()
	n.Double()
	memento := n.CreateMemento() //记录此时是28
	n.Half()  //一半14
	n.ReinstateMemento(memento)  //看备忘录里的值
	fmt.Println(n.value)
}
