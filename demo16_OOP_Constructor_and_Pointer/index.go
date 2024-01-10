package main

import "demo16/employee"

func main() {
	//encapsulation employee and use public function (New) for action to employee struct in other package
	e := employee.Init(
		"Sam",
		"Adolf",
		30,
		20,
	)

	e = employee.Init(
		"lek",
		"Adolf",
		30,
		20,
	)
	e.LeavesRemaining()
}
