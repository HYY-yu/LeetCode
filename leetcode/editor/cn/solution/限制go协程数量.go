package solution

import (
	"fmt"
)

func run() {
	fmt.Println("xx")
}

func main() {
	channel := make(chan int, 10)

	for i := 0; i < 1000; i++ {
		channel <- 1
		go func() {
			defer func() {
				<-channel
			}()
			run()
		}()
	}
}
