package main

import "fmt"

func main() {
	fmt.Println("Decimal to Binary!")

	var number int = 0

	fmt.Print("Please input number: ")
	fmt.Scan(&number)
	fmt.Printf("number = %d, in binary format = %b\n", number, number)

}
