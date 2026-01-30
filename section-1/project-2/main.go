package main

import "fmt"

type UserLevel int

const (
	Basic UserLevel = iota
	Premium
	VIP
)

func (u UserLevel) String() string {
	switch u {
	case Basic:
		return "User level: Basic \n Akses: Limit Akses"
	case Premium:
		return "User Level: Premium \n Akses: Premium Akses"
	case VIP:
		return "User Level: VIP \n Akses: VIP Akses"
	default:
		return "unkown"
	}
}

func main() {
	level := Basic

	fmt.Println(level)
}
