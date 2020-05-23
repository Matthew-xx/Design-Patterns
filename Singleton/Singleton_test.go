package Singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	//p := GetInstance()  //获取单例（实例
	//p.IncrementAge()  //调用单例里面的方法

	wg := sync.WaitGroup{}
	wg.Add(200)

	for i:=0; i<100; i++ {
		go func() {
			defer wg.Done()
			IncrementAge()
		}()
		go func() {
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
	p := GetInstance()
	age := p.GatAge()
	fmt.Println(age)
}
