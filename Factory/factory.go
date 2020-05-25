package Factory

import "fmt"

//工厂生产、客户取物

//接口（提供getfood方法
type Restaurant interface {
	GetFood()
}

//两个饭店（结构体）都生产自己的食物：创建实现接口的实体类。
type Donglaishun struct {

}

func (d *Donglaishun) GetFood()  {
	fmt.Println("东来顺饭菜已做好")
}

type Qingfeng struct {

}

func (q *Qingfeng) GetFood()  {
	fmt.Println("庆丰包子已做好")
}

//创建工厂
func NewRestaurant(name string) Restaurant{
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}

