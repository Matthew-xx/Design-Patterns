package Visitor

import "testing"

func TestElement_Accept(t *testing.T) {
	e := new(Element)
	e.Accept(new(WeiBoVisitor))  //接收微博访问者，accept调用实现了接口的实例
	e.Accept(new(IQIYIVisitor))
}

