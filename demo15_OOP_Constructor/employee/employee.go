package employee

import (
	"fmt"
)

type employee struct { //protect employee use only this package
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstname string, lastname string, totalLeaves int, leavesToken int) employee {
	e := employee{
		FirstName:   firstname,
		LastName:    lastname,
		TotalLeaves: totalLeaves,
		LeavesTaken: leavesToken}
	return e

}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
