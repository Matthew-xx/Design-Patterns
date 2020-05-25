package Observer

import (
	"sync"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	//for x:= range Fib(10){
	//	fmt.Println(x)
	//}

	n := eventSubject{Observers:sync.Map{}}  //如关注的微博主更新微博
	obs1 := eventObserver{ID:1,Time:time.Now()}
	obs2 := eventObserver{ID:2,Time:time.Now()}

	n.AddListener(obs1)
	n.AddListener(obs2)

	for x:=range Fib(10){
		n.Notify(Event{Data:x})  //关注的微博循环发通知
	}
}
