package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)
	var superman string
	superman = "I am superman "
	fmt.Println(superman)
	thor := "I am thor"
	fmt.Println(thor)

	//thorRating := 3.
	//fmt.Printf("%v, %T", thorRating, thorRating)

	var Ironman, CatAmerica string = "I am ironman", "I am catAmerica"
	fmt.Println(Ironman, CatAmerica)

	var defaultValue string
	fmt.Println(defaultValue)

	var (
		spiderman = "I am spiderman "
		age       = 18
		power     = "web slings, spidy sense"
		antman    = "I am antman"
	)
	fmt.Println(spiderman, age, power, antman)
}
