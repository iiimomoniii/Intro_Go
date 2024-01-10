package main

import (
	"fmt"
	"time"
)

func main(){
	channel();
}

func channel(){
	ch := make(chan int, 1); // 1 = channel
	//send 1 to channel
	ch <- 1
	fmt.Println("Step 1")
	fmt.Println(<-ch)
	//send 2 to channel
	ch <- 2
	fmt.Println("Step 2")
	fmt.Println(<-ch)

	time.Sleep(1*time.Second)
}