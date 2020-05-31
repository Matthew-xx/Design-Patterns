package main

import (
	"fmt"
	"sort"
)

//
//import (
//	"../Design-Patterns/structural/Decorator"
//	"fmt"
//	"log"
//	"os"
//	"time"
//)
//func main()  {
//	//fmt.Println(Decorator.Pi(100))
//	f := Decorator.WrapLogger(Decorator.Pi,log.New(os.Stdout,"test ",1))
//	f(10000)
//
//	const Layout = "2006-01-02 15:04:05"//时间常量
//	//loc, _ := time.LoadLocation("Asia/Shanghai")
//	xtime := time.Now().Format("2006-01-02 15:04:05")
//	mtime,_ := time.ParseInLocation(Layout,xtime, time.Local)
//	ntime := time.Date(mtime.Year(),mtime.Month(),mtime.Day(),mtime.Hour()+1,mtime.Minute(),mtime.Second(),mtime.Nanosecond(),time.Local)
//
//	fmt.Println(mtime)
//	fmt.Println(ntime)
//
//
//}



type Person struct {
	Name string
	Age int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}