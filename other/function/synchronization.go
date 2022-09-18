package function

import (
	"fmt"
	"sync"
	"time"
)

var l sync.Mutex
var count int = 0

func add(){
	l.Lock()
	count++
	l.Unlock()
}

func Test(){
	for i:=1; i <= 10000; i++{
		go add()
	}
	time.Sleep(2*time.Second)
	fmt.Println(count)
}

func Test2(){
	var once sync.Once
	/// 打印结果为1 因为只执行一次
	once.Do(add)
	once.Do(add)
	once.Do(add)
	fmt.Println(count)
}

/// Add添加要等待几个 Wait开始等待 Done表示完成的+1 当Done的数量达到一开始Add时
/// 就继续执行
var a sync.WaitGroup

func Test_WaitGroup(){
	
	a.Add(2)

	go wait()
	time.Sleep(2*time.Second)
	go wait()

	a.Wait()
	fmt.Println("fini")
}

func wait(){
	time.Sleep(2*time.Second)
	fmt.Println("fiber fini")
	a.Done()
}