package main

import (
	"fmt"
)

func main(){
	//fnArray();
	//fnSliceArrayAdd()
	//fnSliceArrayRemove()
	fnRemoveIndex();
}

//Array
func fnArray(){
	//full
	var array1 []int = []int {1,2,3,4};
	fmt.Println(array1[0]);
	//short
	var array2 = []int {1,2,3,4};
	fmt.Println(array2[0]);
	//assign value after decared
	var array3 [3]string;
	array3[0],array3[1],array3[2] = "android","iOS","react";
	fmt.Println(array3[1]);

	//for print value in array
	for _, item := range array1 {
		fmt.Println(item);
	}

	for _, item := range array3 {
		fmt.Println(item);
	}
}

//Slice Array Add
func fnSliceArrayAdd(){
	var numbers = make([] int, 3, 5); // 3 = len , 5 = cap
	numbers = append(numbers,2);
	numbers = append(numbers,4);
	numbers = append(numbers,6);
	numbers = append(numbers,8); // over size (cap 5 + 5 = cap 10 )
	viewData(numbers);

	var number2s []int;
	number2s = append(number2s, 1);
	number2s = append(number2s, 3);
	number2s = append(number2s, 5);
	viewData(number2s);
}


//Slice Array Remove
func fnSliceArrayRemove(){
	var number3s = []int {1,3,5,7,9}
	viewData(number3s);

	//remove by leading remove
	number3s = number3s[2: len(number3s)] // [Quantity to remove : len()]
	viewData(number3s); // 5 7 9 

	//remove by trailing remove
	number3s = number3s[0: len(number3s)-1] // remove last index of array
	viewData(number3s) // 5 7
}

//Remove specific Index
func fnRemoveIndex() {
	var number4s = []int {1,3,5,7,9};
	viewData(number4s);
	number4s = removeIndex(number4s, 2); // remove index = 2
	viewData(number4s);
}


func removeIndex(s []int, index int)[]int{ // s = original slice
	return append(s[:index], s[index+1:]...) //  
	//how to process
	//1. s[:index] => (sub slice 0-> index (indentify) )
	//2. s[index+1:] => slice after index (indentify) to last index of array
	//3. append 

}

//Center function for View Data
func viewData(x []int){
	//len show value in array size
	//cap show maximun capacity
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x); 
}