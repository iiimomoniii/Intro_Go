package employee

import (
	"fmt"
)

var employeeInstance *employee //use *employee for validate empty data

type employee struct { //protect employee use only this package
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func Init(firstname string, lastname string, totalLeaves int, leavesToken int) *employee {
	//validate employeeInstance is created
	if employeeInstance == nil {
		//if employeeInstance not exiting we will create
		employeeInstance = &employee{
			FirstName:   firstname,
			LastName:    lastname,
			TotalLeaves: totalLeaves,
			LeavesTaken: leavesToken}
	}
	
	return employeeInstance

}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
