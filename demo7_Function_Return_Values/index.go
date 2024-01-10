package main

import "fmt"

func main() {
	fmt.Println(sum(2,3));

	//cal from value
	const a,b = 5,7;
	fmt.Println(sum(a,b));

	//print decimal format (%d) and cal values
	const c,d = 9,16;
	fmt.Printf("%d+%d=%d\n",c,d,sum(c,d));

	//print return mutiple values
	var x,y int = swap(a,b);
	//short form
	//x,y = swap(a,b);
	fmt.Printf("%d,%d => %d,%d\n", a,b, x,y)

	//print return mutiple values and mutiple types
	_w,_z,title := swap3(10,20);
	fmt.Printf("%d,%d => %d,%d,%s\n", 10,20, _w,_z,title)
}

// function calculate
func sum(num1 int, num2 int) int {
	return num1 + num2
}

// function calculate return mutiple values
func swap(a int, b int) (int, int) {
	return b,a;
}

// function calculate return mutiple values with naming
func swap2(a int, b int) (x,y int) {
	x = b
	y = a
	return
}

// function calculate return mutiple values and mutiple types with naming
func swap3(a int, b int) (x,y int, title string) {
	x = b
	y = a
	title = "Result"
	return
}