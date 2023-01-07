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


func arrayThings()[]string{
	var a []string
	// a[0]="sisi"
	a=append(a,"mormon",)
	a=append(a,"pp",)
return a
}

func main() {
	var count uint=0
	for {
count++
fmt.Println(count)
if count%2!=0 {
	continue
}
		fmt.Printf("Hello, %v %T!\n",arrayThings(),name1)
		if count>=20 {
			break
		}
	}


	// fmt.Println(quote.Go())
}