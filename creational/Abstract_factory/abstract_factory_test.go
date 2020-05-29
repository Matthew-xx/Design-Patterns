package Abstract_factory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factory := NewSimpleLunchFactory()
	food := factory.CreateFood()
	food.Cook()

	vagetable := factory.CreateVagetable()
	vagetable.Cook()
}
