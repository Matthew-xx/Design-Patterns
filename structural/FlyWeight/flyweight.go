package FlyWeight

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name:name}
}

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

//初始化pool
func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool:make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight  {
	weight,ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name)  //pool里面有就直接返回，没有就new一个进pool里面再返回
		f.pool[name] = weight
	}
	return weight
}

