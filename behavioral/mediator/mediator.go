package mediator

import "fmt"

type Mediator interface {
	Communicate(who string)
}

type WildStallion interface {
	SetMediator(mediator Mediator)
}

type Bill struct {
	mediator Mediator
}

func (b *Bill) SetMediator(mediator Mediator)  {
	b.mediator = mediator
}

func (b *Bill) Resond()  {
	fmt.Println("bill what ...")
	b.mediator.Communicate("bill")
}

type Ted struct {
	mediator Mediator
}

func (t *Ted) Talk()  {
	fmt.Println("ted:bill...")
	t.mediator.Communicate("ted")
}

func (t *Ted) SetMediator(mediator Mediator)   {
	t.mediator = mediator
}

func (t *Ted) Respond()  {
	fmt.Println("ted:how are you today?")
}

type ConcreteMediator struct {
	Bill
	Ted
}

func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

func (m *ConcreteMediator)Communicate(who string)  {
	if who == "Ted" {
		m.Bill.Resond()
	}else {
		m.Ted.Respond()
	}
}


