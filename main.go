package main

import "fmt"

type User struct {
	name    string
	surname string
	age     int
}

func main() {
	user1 := User{name: "Serdar", surname: "Cetger", age: 18}
	user2 := User{name: "Serdar", surname: "Cetger", age: 25}
	//we dont need temp variable to exchange :*(
	user1.age, user2.age = user2.age, user1.age
	fmt.Println(user1)
	fmt.Println(user2)

}
