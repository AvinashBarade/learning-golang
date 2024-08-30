package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof!"
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is of %d age", p.Name, p.Age)
}

func ptintType(i interface{}) {

	switch v := i.(type) {
	case string:
		fmt.Println("String", v)
	case int:
		fmt.Println("Int", v)

	}
}

func main() {
	var s Speaker
	s = Dog{}
	//d := Dog{}
	fmt.Println(s.Speak())
	ptintType("hello")
	ptintType(45)

	p := Person{
		Name: "avinash",
		Age:  29,
	}
	fmt.Println(p)
}
