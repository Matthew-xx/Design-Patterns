package Strategy

import "fmt"

//定义抽象策略接口（及接口方法
type Strategy interface {
	Execute()
}

//定义策略A
type StrategyA struct {

}
//策略A实现了接口
func (s *StrategyA) Execute()  {
	fmt.Println("a plan executed..")
}

func NewStrategyA() Strategy {
	return &StrategyA{}
}


type StrategyB struct {
	
}

func (s *StrategyB) Execute()  {
	fmt.Println("B PLAN EXECUTED ")
}

func NewStrategyB() Strategy {
	return &StrategyB{}
}
//定义环境类（操作者，选择什么策略
type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy)  {
	c.strategy = strategy
}
//调用策略实现的方法
func (c *Context) Execute()  {
	c.strategy.Execute()
}
