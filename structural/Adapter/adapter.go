package Adapter

import "fmt"

type Target interface {
	Execute()
}

type Adaptee struct {
	
}

func (a *Adaptee) SpecificExecute()  {
	fmt.Println("充电...")
}

//适配器
type Adapter struct {
	*Adaptee  //适配器调用真实的充电
}

func (a *Adapter) Execute()  {
	a.SpecificExecute()
}

