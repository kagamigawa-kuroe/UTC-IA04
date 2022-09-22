package Ex2

import (
	"fmt"
	"sync"
)

var n = 0

var m sync.Mutex

func f() {
	m.Lock()
    n++
	m.Unlock()
}

func Main_Ex2() {
    for i := 0; i < 10000; i++ {
        go f()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("n:", n)
}
