package Template

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}

type Worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}

func (w *Worker) Daily()  {
	w.Work()
	w.GetUp()
	w.Sleep()
}

//coder实现了WorkInterface 的三个方法
//要熟练掌握封装，继承，多态
type Coder struct {
	//称为worker的子类
}

func (c *Coder)GetUp()  {
	fmt.Println("coder getup")
}

func (c *Coder)Work()  {
	fmt.Println("coder work")
}

func (c *Coder)Sleep()  {
	fmt.Println("coder sleep")
}
