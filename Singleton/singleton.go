package Singleton

import "sync"

//构造对象
var(
	p *Pet
	once sync.Once
)

//只初始化一次
func init()  {
	once.Do(func() {
		p=&Pet{}
	})
}

func GetInstance() *Pet {
	return p
}

type Pet struct {
	name string
	age int
	mux sync.Mutex
}

//单例模式一定要是线程安全的(加锁、释放锁）
func (s *Pet) SetName(n string)  {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.name = n
}

//生产环境下读不要加锁（并发情况下
func (p *Pet)GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *Pet) IncrementAge()  {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.age++
}

func (p *Pet) GatAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}

