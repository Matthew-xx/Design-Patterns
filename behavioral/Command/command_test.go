package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	laowang := NewPerson("wang",NewCommand(nil,nil))
	laozhang := NewPerson("zhang",NewCommand(&laowang,laowang.Listen))
	laoma := NewPerson("ma",NewCommand(&laozhang,laozhang.Buy))
	laochen := NewPerson("chen",NewCommand(&laoma,laoma.Cook))
	laoli := NewPerson("li",NewCommand(&laochen,laochen.Wash))
	laoli.Talk()
}

//先调laoli.talk,然后NewCommand(&laochen,laochen.Wash)即命令老陈laochen.wash
//调到最上面laowang.Listen，老王没有可调了NewCommand(nil,nil)，所以在listen那的函数要将execute去掉，否则报错