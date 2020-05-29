package Facade

import "fmt"

type CarModel struct {

}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel()  {
	fmt.Println("carmodel - setmodel")
}

type CarEngine struct {
	
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine()  {
	fmt.Println("carengine - setengine")
}

type CarBody struct {

}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetCarBody()  {
	fmt.Println("carbody - setbody")
}

type CarFacade struct {
	model CarModel
	engine CarEngine
	body CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:CarModel{},
		engine:CarEngine{},
		body:CarBody{},
	}
}

func (c *CarFacade) CreateCompleteCar()  {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetCarBody()
}



