package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	rob := User{"rob", "rob@loc.dev", 34}
	fmt.Printf("%+v\n", rob)
	fmt.Printf("%v\n", rob.Email)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@loc.dev"
	sam.Age = 22
	fmt.Printf("%+v\n", sam)

	var tobby = &User{"tobby", "tobby@loc.dev", 22}
	fmt.Printf("%+v\n", tobby)

}
