package FlyWeight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFlyWeightFactory_GetFlyWeight(t *testing.T) {
	factory := NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong beauty")
	xiang := factory.GetFlyWeight("xing xiang")

	assert.Len(t,factory.pool,2)
	assert.Equal(t,hong,factory.pool["hong beauty"])
	assert.Equal(t,xiang,factory.pool["xing xiang"])
}
