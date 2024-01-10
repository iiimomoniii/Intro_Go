package main

import "demo15/employee"

func main() {
	//encapsulation employee and use public function (New) for action to employee struct in other package
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20,
	)
	e.LeavesRemaining()
}
