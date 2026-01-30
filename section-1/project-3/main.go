package main

import "fmt"

type UserLevel int

const (
	Basic UserLevel = iota
	Premium
	VIP
)

var levelData = []struct {
	Name  string
	Akses string
}{
	{"Basic", "Limit Akses"},
	{"Premium", "Premium Akses"},
	{"VIP", "VIP Akses"},
}

func (u UserLevel) String() string {
	if int(u) < 0 || int(u) >= len(levelData) {
		return "Unknown"
	}

	return fmt.Sprintf(
		"User Level: %s\nAkses: %s",
		levelData[u].Name,
		levelData[u].Akses,
	)
}

func main(){
	userAccess :=Premium

	fmt.Println(userAccess)
}