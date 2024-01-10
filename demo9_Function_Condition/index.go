package main

import (
	"fmt"
)

func main(){
	// fnIfAndElse();
	// fnIfAndElse2();
	// fnValidateResult();
	fnSwitchCase();
}

//If else
func fnIfAndElse(){
	someValue := 11;
	if (someValue == 10) {
		fmt.Println("equal 10");
	} else {
		fmt.Println("not equal 10");
	}
}

func fnIfAndElse2(){
	someValue := 11;
	if (someValue == 10 || someValue < 2) {
		fmt.Println("someValue == 10 || someValue < 2");
	} else {
		fmt.Println("not someValue == 10 || someValue < 2");
	}
}

//Validate value from other function
func fnValidateResult(){
	if result := doSomething(); result == "Hello" {
		fmt.Println("Yeah Hello")
	 } else {
		fmt.Println("Oop")
	 }
}

func doSomething() string {
	//do something
	return "Hello";
}

//Swith Case
func fnSwitchCase(){
	index := 2;
	switch index {
		case 1 : 
			fmt.Println("1"); 
			break;
		case 2 : 
			fmt.Println("2"); 
			break;
		case 3 : 
			fmt.Println("3"); 
			break;
		default : 
			fmt.Println("something else"); 
			break;
	}
}