package Ex1

import (
	"fmt"
	"time"
)

func compte(n int){
	for i:=0; i<=n; i++{
		fmt.Println(i)
	}
}

func compteMsg(n int, msg string){
	for i:=0; i<=n; i++{
		fmt.Println(msg)
	}
}

func compteMsgFromTo(start int, end int, msg string){
	for i:=start; i<end; i++{
		fmt.Println(msg)
	}
}

func Test_Compte(){
	go compte(10)
	go compteMsg(10,"test")
	go compteMsgFromTo(3,13,"test2")
}

var count int = 0

func Test_Add(){
	for i:=0; i<=9; i++{
		go func() {
			for i:=0; i<=999; i++{
				count++;
			}
		}()
	}
	time.Sleep(time.Second*3)
	fmt.Println(count)
}