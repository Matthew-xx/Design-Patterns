package main

import (
	"../设计模式/Decorator"
	"fmt"
	"log"
	"os"
	"time"
)
func main()  {
	//fmt.Println(Decorator.Pi(100))
	f := Decorator.WrapLogger(Decorator.Pi,log.New(os.Stdout,"test ",1))
	f(10000)

	const Layout = "2006-01-02 15:04:05"//时间常量
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	xtime := time.Now().Format("2006-01-02 15:04:05")
	mtime,_ := time.ParseInLocation(Layout,xtime, time.Local)
	ntime := time.Date(mtime.Year(),mtime.Month(),mtime.Day(),mtime.Hour()+1,mtime.Minute(),mtime.Second(),mtime.Nanosecond(),time.Local)

	fmt.Println(mtime)
	fmt.Println(ntime)


}
