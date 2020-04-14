package main

import (
	"fmt"
)

func main() {
	start := 1
	things := []string{"ram", "sam", "hari", "aashish"}
	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])

	}

	for i := range things {
		fmt.Println(things[i])
	}

	for start < 100 {
		start += start
		if start == 32 {
			break
		}
		fmt.Println("Start is now:", start)
	}

}
