package main

import (
	buddy "GoPractice/codewars"
	fendi "GoPractice/fen"
	ha "GoPractice/hh"
	"fmt"
)

// import "si"

// import "rsc.io/quote"

var name1 = "Mearnsk"

func getName() string {
	name1 := "Mein nama ist Ra "
	name1 = "89ifenti"
	return name1
}

var a []string

func arrayThings(thing string) []string {
	// a[0]="sisi"
	// a = append(a, "mormon", thing)
	a = append(a, thing)
	return a
}
func main() {
	fmt.Println("hiii")
	ha.CallMan()
	count := 0
	var txt1 string
	for {

		fmt.Println("Current txt is", txt1, "\nEnter another")
		fmt.Scan(&txt1)

		fendi.Si(getName())
		fmt.Println(buddy.SumDivisors(75))
		// fmt.Println(buddy.Buddy(1071625,1103735))
		fmt.Printf("Hello, %v!\n\n=============\n\n", arrayThings(txt1))
		count++
		if count >= 2 {
			break
		}
	}
}
