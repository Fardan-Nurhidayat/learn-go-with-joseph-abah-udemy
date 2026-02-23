package main

import "fmt"

// Go sendiri tidak memiliki class, namun kita bisa menggunakan struct untuk membuat object

// ini adalah receiver dengan struct
type User struct {
	name string
	age  int8
}

// ini adalah receiver dengan type
type Grade int8

var URL string

// pendekalrasian type sebagai interface
// interface adalah kumpulan method signature
// Di Go tidak perlu deklarasi implements , extend , asalkan namanya sesuai langsung dianggap implementasi
type Protocal interface {
	request(URLEndpoint string)
	response() string
}
type HTTPProtocol struct {
	URL string
}

// Method adalah function yang memiliki receiver
// Receiver adalah parameter yang dimiliki oleh method
// Receiver bisa berupa value atau pointer
// Receiver bisa berupa struct atau interface
func (u User) SayHello() {
	fmt.Println("Hello", u.name)
}

// ini adalah method dengan receiver type
func (g Grade) GetGrade() string {
	switch g {
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "F"
	default:
		return "Unknown"
	}
}

func (HTTP *HTTPProtocol) request(URLEndpoint string) {
	URL = fmt.Sprintf("%s/%s", HTTP.URL, URLEndpoint)
	HTTP.URL = URL
}
func (http HTTPProtocol) response() string {
	return URL
}

func main() {
	user := User{
		name: "John",
		age:  30,
	}

	user.SayHello()

	grade := Grade(1)
	fmt.Println(grade.GetGrade())

	HTTPP := HTTPProtocol{URL: "https://www.golinuxcloud.com"}
	HTTPP.request("go-methods")
	res := HTTPP.response()
	fmt.Println(res)
	fmt.Println(HTTPP)
}
