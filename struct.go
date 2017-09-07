package main

import "fmt"

// struct 的使用

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("Your name is %v and  age is %v", u.Name, u.Age)
}

func main() {
	u := User{}
	fmt.Print("Input Your Name > ")
	fmt.Scanf("%s\n", &u.Name)
	fmt.Print("Input Your Age > ")
	fmt.Scanf("%d\n", &u.Age)
	fmt.Println(u.String())
}
