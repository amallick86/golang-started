package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//frontend
	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name ")
	name, _ = reader.ReadString('\n')

	//take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate out Dosa center between 1 and 5:")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//backend
	fmt.Printf("Hello %v, \n Thanks for rating our Dosa center with %v star rating . \n\n Your rating was recorded in our system at %v \n\n", name, mynumRating, time.Now().Format(time.Stamp))
	if mynumRating == 5 {
		fmt.Println("Bonus for teamfor 5 star service")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("we are always improving")
	} else if mynumRating < 3 {
		fmt.Println("Need serious wor on our side")
	}

}
