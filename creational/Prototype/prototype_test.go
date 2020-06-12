package Prototype

import (
	"github.com/stretchr/testify/assert"  //断言
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "浪"
	p := ConcretePrototype{name:name}
	newProto := p.Clone()
	assert.Equal(t,name,newProto.Name())  //比较两个名字是否一致
}


