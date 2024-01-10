package main

import (
	"fmt"
)

func main(){
	// fnMap();
	// fnMap2();
	fnMap3();
}

//Normal Map function
func fnMap(){
	var numbers = map[string]int{"one":1,"two":2,"three":3}
	//view value by key
	fmt.Println("",numbers["two"])
}

//Map Dynamic
func fnMap2(){
	var colors = make(map[string]string)
	colors["red"] = "#f00";
	colors["blue"] = "#00f";
	colors["green"] = "#0f0";
	fmt.Println("",colors);
	fmt.Println("",colors["green"]);
}

//Nested Map
func fnMap3(){
	var courses = make(map[string]map[string]int)
	courses["Android"] = make(map[string]int)
	courses["Android"]["Price"] = 100;
	courses["Android"]["Code"] = 123456;

	courses["iOS"] = make(map[string]int)
	courses["iOS"]["Price"] = 200;
	courses["iOS"]["Code"] = 987654321;
	fmt.Println(courses)
	fmt.Println(courses["iOS"]["Code"])
}