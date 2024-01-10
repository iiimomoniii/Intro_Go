package main

import (
	"fmt"
)

func main(){
	// fnFor();
	// fnWhile();
	// fnWhile2();
	// fnRange();
	// fnRange2();
	fnBreak();
}

// function For Loop
func fnFor(){
	for index := 0; (index <= 10 ); index++{
		fmt.Printf("Index %d\n", index);
	}
}

// function While
func fnWhile(){
	index := 0;
	for index < 5 {
		//1->5
		index ++;
		fmt.Printf("Index %d\n", index);
	}
}

// function While
func fnWhile2(){
	index := 0;
	for index < 5 {
		//0->4
		fmt.Printf("Index %d\n", index);
		index ++;
	}
}

//function For Range (For Each)
func fnRange(){
	courses := []string {"Android", "iOS", "React"};

	//index = seq
	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item);
	}
}

func fnRange2(){
	courses := []string {"Android", "iOS", "React"};

	//ignore index
	for _, item := range courses {
		fmt.Printf("%s\n", item);
	}
}

//Assign Break in Function
func fnBreak(){
	index := 0;
	for true {
		// 0 -> 5 break
		if index > 5 {
			break
		}
		fmt.Printf("While-Index %d\n", index);
		index++;
	}
}