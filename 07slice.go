package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"maleta", "ropa", "reloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "iroman"
	heros[2] = "spiderman"
	fmt.Println(heros)
	heros = append(heros, "deadpool")
	heros = append(heros, "batman")
	fmt.Println(heros)
	fmt.Println(cap(heros))

	myints := []int{4, 7, 3, 2, 8}

	isSorted := sort.IntsAreSorted(myints)
	fmt.Println("Are ints sorted:", isSorted)
	sort.Ints(myints)
	isSorteds := sort.IntsAreSorted(myints)
	fmt.Println("Are ints sorted:", isSorteds)

	fmt.Println(myints)
}
