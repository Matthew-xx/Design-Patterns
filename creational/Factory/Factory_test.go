package Factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	//NewRestaurant("d")新建饭店，通过标识"d"调用不同工厂的方法
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()
}
