package Visitor

import "fmt"

//定义了对每个 Element 访问的行为，它的参数就是被访问的元素
type IVisitor interface {
	Visit()
}

//具体的访问者，它需要给出对每一个元素类访问时所产生的具体行为。
type WeiBoVisitor struct {
	
}

func (w WeiBoVisitor) Visit()  {
	fmt.Println("visit weibo")
}

type IQIYIVisitor struct {

}

func (i IQIYIVisitor) Visit()  {
	fmt.Println("visit iqiyi")
}

//元素接口或者抽象类，它定义了一个接受访问者（accept）的方法，其意义是指每一个元素都要可以被访问者访问
type IElement interface {
	Accept(visitor IVisitor)
}

//具体的元素类，它提供接受访问的具体实现，而这个具体的实现，通常情况下是使用访问者提供的访问该元素类的方法(visit)。
type Element struct {

}

func (e Element) Accept(v IVisitor) {
	v.Visit()
}
