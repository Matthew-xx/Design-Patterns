package Proxy

import (
	"fmt"
	"strconv"
)

type ITask interface {
	RentHouse(desc string,price int)
}

type Task struct {

}

func (t *Task) RentHouse(desc string,price int)  {
	fmt.Println(fmt.Sprintf("租房地址%s,价格%s",desc,strconv.Itoa(price)))
}

//代理
type AgentTask struct {
	task *Task
}

func NewAgentTask() *AgentTask {
	return &AgentTask{task:&Task{}}
}

func (t *AgentTask) RentHouse(desc string,price int)  {
	t.task.RentHouse(desc,price)
}
