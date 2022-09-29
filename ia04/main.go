package main

import (
	"fmt"
	"ia04/agt"
)

func TestAgt()  {
	/// create
	a := make([]agt.Alternative,0)
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

}

func main(){
	TestAgt()
}