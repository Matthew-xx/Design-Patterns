package Builder

import (
	"fmt"
	"testing"
)

func TestConcreateBuilder_GetResult(t *testing.T) {
	builder := NewConcreateBuilder()
	director := NewDirector(&builder)
	director.Construct()
	result := builder.GetResult()
	fmt.Println(result.Built)
}
