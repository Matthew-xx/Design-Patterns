package Observer

import (
	"fmt"
	"sync"
	"time"
)
/*
Subject类是主题，它把所有对观察者对象的引用文件存在了一个聚集里，每个主题都可以有任何数量的观察者。
抽象主题提供了一个接口，可以增加和删除观察者对象；
Observer类是抽象观察者，为所有的具体观察者定义一个接口，在得到主题的通知时更新自己；
ConcreteSubject类是具体主题，将有关状态存入具体观察者对象，在具体主题内部状态改变时，给所有登记过的观察者发出通知；
ConcreteObserver是具体观察者，实现抽象观察者角色所要求的更新接口，以便使本身的状态与主题的状态相协同。
 */


//数据
type Event struct {
	Data int
}
//观察者(观察者一般是一个接口，每一个实现该接口的实现类都是具体观察者。
type Observer interface {
	NotifyCallback(event Event)  //通知的方法（主题变更后通知
}
//主题（接口
type Subject interface {
	AddListener(observer Observer)  //添加观察者
	RemoveListener(observer Observer) 
	Notify(event Event)  //通知
}

//对应观察者
type eventObserver struct {
	ID int
	Time time.Time
} 

//对应主题(继承Subject类，在这里实现具体业务，在具体项目中，该类会有很多变种。
type eventSubject struct {
	Observers sync.Map
}

//eventObserver类实现了Observer接口的方法
func (e eventObserver)NotifyCallback(event Event)  {
	fmt.Printf("recieved:%d after %v\n",event.Data,time.Since(e.Time))
}

func (e *eventSubject) AddListener(obs Observer)  {
	e.Observers.Store(obs, struct {}{})
}

func (e *eventSubject)RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}
//使用指针，减少拷贝减少使用内存空间，且可改变原始值
func (e *eventSubject)Notify(event Event)  {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i,j:=0,1;i<n;i,j=i+j,i {
			out <- i
		}
	}()
	return out
}
