package Iterator

import (
	"fmt"
	"testing"
)

func TestArrayIterator(t *testing.T) {
	array := []interface{}{1,2,4,5,7,9}
	a := 0
	iterator := ArrayIterator{array:array,index:&a}
	for it:=iterator;iterator.HasNext() ;iterator.Next()  {
		//取到索引、值，判断后续是否还有东西，若有则取出来返回给迭代器
		index,value := it.Index(),it.Value().(int)
		for value != array[*index] {
			fmt.Println("error")
		}
		fmt.Println(*index,value)
	}
}
