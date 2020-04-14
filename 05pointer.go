package main

import "fmt"

func main() {
	// var p *int

	// if p != nil {
	// 	fmt.Println("p is having a value:", *p)
	// } else {
	// 	fmt.Println("watchout for nil values")
	// }
	var lifebooster float64 = 99.2
	lifeboosterRef := &lifebooster
	ref := &lifeboosterRef
	fmt.Println(lifebooster)
	fmt.Println(lifeboosterRef)
	fmt.Println(*lifeboosterRef)
	fmt.Println(ref)

}
