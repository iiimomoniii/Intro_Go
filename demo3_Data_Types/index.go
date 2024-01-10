package main

import "fmt"
func main() {
	fmt.Println("Begin")
	//Explicit declaration
	//การระบุ type ของตัวแปร เป็น static type เปลี่ยน type ที่หลังไม่ได้
	var tmp1 int = 0;
	var tmp2 string = "hello";
	var tmp3 bool = true;
	const tmp4 int = 0;
	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)

	//Implicit Declaration
	//การไม่ระบุ type ของตัวแปร
	tmp5 := 0;
	tmp6 := "codemobile";
	tmp7 := false;
	fmt.Println(tmp4)
	fmt.Println(tmp5)
	fmt.Println(tmp6)
	fmt.Println(tmp7)
}
