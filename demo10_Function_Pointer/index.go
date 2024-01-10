package main

import (
	"fmt"
)

//Point index of value for Perfomance (*var = pointer)
func main(){
	msg := "some message";
	var msgPointer *string = &msg;
	fmt.Println(msg);//decimal number
	fmt.Println(*msgPointer); //hexadecimal

	changeMessage(&msg, "New Message 1");
	fmt.Println(msg);

	changeMessage(msgPointer, "New Message 2");
	fmt.Println(msg);
}

//
func changeMessage(aPointer *string, newMessage string){
	*aPointer = newMessage;
}