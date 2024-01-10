package main

import "fmt"

//0. create global variable
var externalStr string = "print from external function";

func main() {
	//2.call external function
	external()
}

//1.create external function
func external() {
	//3.print global variable
	fmt.Println(externalStr);
}
