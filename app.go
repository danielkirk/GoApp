package main

import "fmt"

const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessDenied() {
	fmt.Println("Access Denied")
}

func accessGranted() {
	fmt.Println("Access Granted")
}

func main() {
	today, role := Tuesday, Guest
	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && (today == Saturday || today == Sunday) {
		accessGranted()
	} else if role == Member && !(today == Saturday || today == Sunday) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}
}
