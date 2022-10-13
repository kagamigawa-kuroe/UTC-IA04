package main

import (
	"fmt"
	"ia04/agt"
	"ia04/comsoc"
	"sync"
)

func TestAgt()  {
	/// create
	a := make([]comsoc.Alternative,0)
	agant := agt.NewAgent(1,"test",a)

	/// string
	fmt.Println(agant.String())

	/// prefer
	agant.Prefers(9,8)
	fmt.Println(agant.String())
	agant.Prefers(8,7)
	agant.Prefers(6,5)
	fmt.Println(agant.String())

	/// clone
	agt2 := agant.Clone()
	fmt.Println(agt2.String())

	/// Equal/DeepEqual
	fmt.Println(agt2.Equal(agant))
	fmt.Println(agt2.DeepEqual(agant))
}

func TestCom() {
	prefs := [][]comsoc.Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, _ := comsoc.BordaSWF(prefs)
	fmt.Println(res)
}

func test_demo(){
	var a sync.Mutex
	ag := agt.NewBureau(a,"localhost:8005",100,"server",nil,nil)
	ag.Alts = make([]comsoc.Alternative,0)
	ag.P = make([][]comsoc.Alternative,0)
	go ag.Start()

	fmt.Scanln()

	agent := agt.NewAgent(101,"test",nil)
	agent.Prefs = make([]comsoc.Alternative,0)
	agent.Prefers(1,2)
	agent.Prefers(3,5)
	c := agt.NewVoteur(*agent,"http://localhost:8005")
	fmt.Println(c.Prefs)
	fmt.Scanln()
	c.StartVote()

	fmt.Scanln()

	fmt.Println(c.RequestAnswer())
}

func main(){
	test_demo()
}