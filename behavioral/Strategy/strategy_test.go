package Strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	strategyA := NewStrategyA()  //构造A的结构体实例
	C := NewContext()
	C.SetStrategy(strategyA) //将A注入到策略模式的接口里
	C.Execute()  //策略模型执行注入进来的A

	strategyB := NewStrategyB()
	C.SetStrategy(strategyB)
	C.Execute()
}

