package Ex3

import (
	"fmt"
	"time"
)

func BonneAnnee1(){
	for i := 5; i >= 2; i-- {
		fmt.Print(i)
		fmt.Printf("%c",',')
		time.Sleep(time.Second)
	}
	print("1")
	time.Sleep(time.Second)
	print("... Bonne année\n")
}

// a := time.After(5*time.Second)
// <-a
func BonneAnnee2(){
	for i := 5; i >= 2; i-- {
		fmt.Print(i)
		fmt.Printf("%c",',')
		a := time.After(time.Second)
		<-a
	}
	print("1")
	a := time.After(time.Second)
	<-a
	print("... Bonne année\n")
}

func BonneAnnee3(){
	i := 5
	fmt.Print(i)
	for range time.Tick(time.Second) {
		if(i == 1){
			break;
		}
		i--
		fmt.Printf("%c",',')
		fmt.Print(i)
	}
	fmt.Printf("%c",',')
	fmt.Print(1)
	print("... Bonne année\n")
}

func Test(){
	BonneAnnee3()
}