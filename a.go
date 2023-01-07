package main

import "fmt"

// import "si"

// import "rsc.io/quote"


var name1="Mearnsk"

func getName ()string{
	name1 := "Mein nama ist Ra "
	name1="89"
	return name1
}
func main() {
	var name2 string
	fmt.Scan(&name2)
    fmt.Printf("Hello, R%vAdi%T!\n",name1,name1)
    fmt.Println("Hello, Radi!",name2)
    fmt.Println("Hello, World! "+name1,getName(),name1)

	// fmt.Println(quote.Go())
}