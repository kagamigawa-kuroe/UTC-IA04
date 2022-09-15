package mypackage

import "fmt"
import "math/rand"
import "time"

func Fill(sl []int){
	rand.Seed(time.Now().UnixNano());
	for i:=0; i<len(sl);i++{
		sl[i] = rand.Intn(100)
		fmt.Println(sl[i]);
	}
}
