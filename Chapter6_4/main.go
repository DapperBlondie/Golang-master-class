package main
// using sync.Mutex var for lock/unlocking the shared resources
import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var shres int = 200

func main()  {

	go func() {

		for j := 0;j < 10;j++ {

			mutex.Lock()
			shres += j
			mutex.Unlock()
		}

		return
	}()

	go func() {

		for j := 0;j < 10;j++ {

			mutex.Lock()
			shres -= j
			mutex.Unlock()
		}


		return
	}()

	time.Sleep(time.Second)
	fmt.Println("shared resource : ", shres)

	return
}
