package Ex4

import (
	"fmt"
	"time"
)

func Fill(tab []int, v int){
	for i:=0; i < len(tab); i++{
		tab[i] = v
	}
}

func Fill_Avance(tab []int, v int){
	a := func (tab []int, v int, start int, end int)  {
		for i:=start; i<=end; i++{
			tab[i] = v
		}
	}

	size := len(tab)/10000
	for i := 0; i < size; i++ {
		go a(tab,v,i*10000,i*10000+9999)
	}
	go a(tab,v,size*10000,len(tab)-1)

}

func Equal(tab1 []int, tab2 []int) bool{
	if(len(tab1)!=len(tab2)){
		return false;
	}
	for i:=0; i<len(tab1); i++{
		if(tab1[i]!=tab2[i]){
			return false;
		}
	}
	return true;
}

var l chan bool = make(chan bool)
func Equal_Avance(tab1 []int, tab2 []int) bool{
	a := func (tab1 []int, tab2 []int, start int, end int){
		for i:=start; i<=end; i++{
			if(tab1[i]!=tab2[i]){
				l <- false;
			}
		}
		l <- true
	}

	if(len(tab1)!=len(tab2)){
		return false;
	}

	size := len(tab1)/10000
	for i := 0; i < size; i++ {
		go a(tab1,tab2,i*10000,i*10000+9999)
	}
	go a(tab1,tab2,size*10000,len(tab1)-1)

	total := 0;
	ans := true
	for total < size + 1{
		s := <- l
		ans = (ans && s)
		total++;
	}
	return ans
}

func TestEx4(){
	var a []int = make([]int,7000000000)
	starttime := time.Now()
	Fill(a,8)
	fmt.Println(time.Since(starttime))
	starttime = time.Now()
	Fill_Avance(a,8)
	fmt.Println(time.Since(starttime))
}