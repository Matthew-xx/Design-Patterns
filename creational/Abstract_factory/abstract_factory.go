package Abstract_factory

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {

}

func (r *Rise) Cook()  {
	fmt.Println("it is rise")
}

type Tomato struct {

}

func (t *Tomato) Cook()  {
	fmt.Println("it is tomato")
}

type LunchFactory interface {
	CreateFood() Lunch   //返回的应是实例而不是行为，而rise等实现了lunch接口，于是便返回rise等实体对象
	CreateVagetable() Lunch
}

type SimpleLunchFactory struct {

}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVagetable() Lunch{
	return &Tomato{}
}


