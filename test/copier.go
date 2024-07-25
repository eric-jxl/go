package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Role string
	Age  int32
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name      string
	Age       int32
	DoubleAge int32
	SuperRole string
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}

func main() {
	user := User{Name: "Jinzhu", Age: 18, Role: "Admin"}
	employee := Employee{}

	err := copier.Copy(&employee, &user)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", employee)
}
