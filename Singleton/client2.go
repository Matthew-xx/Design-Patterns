package Singleton

func IncrementAge2()  {
	p := GetInstance()  //获取单例（实例
	p.IncrementAge()  //调用单例里面的方法
}
