package fendi

import "fmt"

func Si() {
	namesArr := []string{"pploo", "oggieoo", "sogoo", "tresoo", "wasdoo", "arrgoo"}
	namesArr2 := []string{"ppl", "oggie", "sog", "tres", "wasd", "arrg"}
	// ll:="iggii"
	namesArr[len(namesArr)-1] = "agamauu"
	namesArr[1] = "agama"
	namesArr2 = append(namesArr2, "ff", "iii")
	namesArr2 = append(namesArr2, namesArr...)

	fmt.Printf("\n\nfendi %v \n\n", namesArr2)
}
