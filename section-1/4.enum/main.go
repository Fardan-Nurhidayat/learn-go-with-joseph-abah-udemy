// package main

// import "fmt"

// type StatusPembayaran int

// const (
// 	Pending StatusPembayaran = iota
// 	Approved
// 	Rejected
// )

// type LogLevel int

// const (
// 	LogError LogLevel = iota
// 	LogWarn
// 	LogInfo
// 	LogDebug
// 	LogFatal
// )

// type Roles string

// const (
// 	Admin Roles = "admin"
// 	User  Roles = "user"
// 	Guest Roles = "guest"
// )

// func main() {
// 	status := Approved
// 	role := User
// 	fmt.Println(status)
// 	fmt.Println(role)
// }

// func (s StatusPembayaran) String() string {
// 	switch s {
// 	case Pending:
// 		return "Pending"
// 	case Approved:
// 		return "Approved"
// 	case Rejected:
// 		return "Rejected"
// 	default:
// 		return "Unknown"
// 	}
// }

// func (r Roles) String() string {
// 	switch r {
// 	case Admin:
// 		return "Admin"
// 	case User:
// 		return "Pengguna"
// 	case Guest:
// 		return "Guest"
// 	default:
// 		return "unknown"
// 	}
// }

package main

import "fmt"

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn
	LogInfo
	LogDebug
	LogFatal
)

func main() {

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

}