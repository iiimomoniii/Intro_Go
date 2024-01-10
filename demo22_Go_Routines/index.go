package main

import (
	"fmt"
	"time"
)

func main (){
	Sequential();
	Concurrent();
}

func Sequential(){
	run1();
	run2();
}

func Concurrent(){
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}

//Run By Sequential Process
func run1(){
	for i:= 0; i< 100 ;i++{
		fmt.Println("Run 1 : ", i);
	}
}

func run2(){
	for i:= 0;i<100 ;i++{
		fmt.Println("Run 2 : ", i);
	}
}


