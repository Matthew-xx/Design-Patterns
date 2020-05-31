package Responsibility_Chain

import (
	"fmt"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	wang := NewHandler("laowang",nil,1)
	zhang := NewHandler("laozhang",wang,2)
	r := wang.Handler(1)
	fmt.Println(r)
	r = zhang.Handler(2)
	fmt.Println(r)
}
