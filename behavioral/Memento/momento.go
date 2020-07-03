package Memento

type Momento struct {
	state int
}

//Memento 包含了要被恢复的对象的状态
func NewMemento(value int) *Momento  {
	return &Momento{value}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value:value}
}

func (n *Number) Double()  {
	n.value = 2*n.value
}

func (n *Number) Half()  {
	n.value /= 2
}

func (n *Number) Value() int {
	return n.value
}

//创建并在 Memento 对象中存储状态。
func (n *Number)CreateMemento() *Momento {
	return NewMemento(n.value)
}

//从 Memento 中恢复对象的状态
func (n *Number) ReinstateMemento(memento *Momento)  {
	n.value = memento.state
}
