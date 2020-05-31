package Command

import "fmt"

//接收者的角色，也就是最后执行动作的那个对象
type Person struct {
	name string
	cmd Command
}

//定义一个命令角色。一般是一个接口，为所有的命令对象声明一个接口，规范将要进行的命令操作。
type Command struct {
	person *Person
	method func()
}

//构建具体命令角色。具体命令实现了命令接口，定义了动作和接收者之间的绑定关系
func NewCommand(p *Person,method func()) Command {
	return Command{
		person:p,
		method:method,
	}
}

func (c *Command) Execute()  {
	c.method()
}

//请求者持有一个命令对象，有一个行动方法。它会在某个时间点执行行动方法，但不关心是谁具体执行了这个动作。
func NewPerson(name string,cmd Command) Person {
	return Person{
		name ,
		cmd,
	}
}

func (p *Person) Buy()  {
	fmt.Println(fmt.Printf("%s buy",p.name))
	p.cmd.Execute()
}

func (p *Person) Cook()  {
	fmt.Println(fmt.Printf("%s Cook",p.name))
	p.cmd.Execute()
}

func (p *Person) Wash()  {
	fmt.Println(fmt.Printf("%s Wash",p.name))
	p.cmd.Execute()
}

func (p *Person) Listen()  {
	fmt.Println(fmt.Printf("%s Listen",p.name))
}

func (p *Person) Talk()  {
	fmt.Println(fmt.Printf("%s talk",p.name))
	p.cmd.Execute()
}


