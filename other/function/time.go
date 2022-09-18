package function

import (
	"fmt"
	"time"
)


func Test_base(){
	d := time.Now()
	fmt.Println(d)
}

func Test_pass(){
	d := time.Now()
	time.Sleep(2*time.Second)
	fmt.Println(time.Since(d))
}

func Test_Troncature(){
	d := time.Now()
	fmt.Println(d.Truncate(time.Second))
}

func do_some_thing(){
	fmt.Println(time.Now())
}

func Boocle(){
	/// Tick函数会返回一个信道，每过一定时间就会有数据从信道产生
	for range time.Tick(2 * time.Second ) {
		do_some_thing()
	}
}