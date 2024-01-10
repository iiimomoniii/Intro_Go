package main

import (
	"fmt"
)

func main(){

	//print in last process
	defer fmt.Println("Finish");

	for i := 0; i < 10 ; i++ {
		defer printFinish(i)// 9->0
	}

	for i := 0; i < 10 ; i++ {
		fmt.Println("",i); // 0->9
	}
}

func printFinish(i int){
	fmt.Println("Finish",i);
}