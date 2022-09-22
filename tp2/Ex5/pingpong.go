package Ex5

import (
	"fmt"
	"sync"
	"time"
)

var l sync.Mutex

type Agent interface {
	Start()
}

type PingAgent struct {
	ID string
	c  chan string
}

type PongAgent struct {
	ID string
	c  chan string
}

func (p *PingAgent) start(){
	for {
		l.Lock()
		p.c <- "ping"
		v := <- p.c
		fmt.Println("machine"+p.ID+" get "+v)
		l.Unlock()
		time.Sleep(time.Second)
	}
}

func (p *PongAgent) start(){
	for {
		v := <- p.c
		if v == "ping"{
			fmt.Println("machine"+p.ID+" get ping")
			p.c <- "pong"
		}
	}
}

func TestEx5(){
	c := make(chan string,10)
	var ping PingAgent = PingAgent{"123",c}
	var ping2 PingAgent = PingAgent{"125",c}
	var pong PongAgent = PongAgent{"124",c}
	go pong.start()
	go ping.start()
	go ping2.start()
	time.Sleep(100*time.Second)
}
