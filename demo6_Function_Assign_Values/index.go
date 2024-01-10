package main

import "fmt";

func main(){
	fn1();
	fn2("hello", 999);
}

//function non pass argument
func fn1(){
	fmt.Println("codemobiles");
}

//function pass argument
func fn2(title string, version int){
	fmt.Println(title, version);
}