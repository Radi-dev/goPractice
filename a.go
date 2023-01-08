package main

import (
	fendi "GoPractice/fen"
	"fmt"
)

// import "si"

// import "rsc.io/quote"

var name1 = "Mearnsk"

func getName() string {
	name1 := "Mein nama ist Ra "
	name1 = "89"
	return name1
}

var a []string

func arrayThings(thing string) []string {
	// a[0]="sisi"
	a = append(a, "mormon", thing)
	a = append(a, "pp", thing)
	return a
}

func main() {
	var count uint = 0

	for index, v := range arrayThings("hent") {
		fmt.Println(index, v)
	}
	for {
		count++
		fmt.Println(count)
		if count%2 != 0 {
			fendi.Si()
			continue
		}
		fmt.Printf("Hello, %v %T!\n", arrayThings("ooi"), name1)
		if count >= 20 {
			break
		}
	}

	// fmt.Println(quote.Go())
}
