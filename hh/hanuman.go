package ha

import "fmt"

func CallMan() {
	// var human map[string]string
	genders := make(map[string]string)
	genders["male"] = "Male"
	genders["female"] = "Memale"
	var man = human{
		description: "A hooman with the capability of breathing", gender: genders["male"], height: 6,
	}
	var human = make(map[string]string)

	human["SunGod"] = "Ra"
	fmt.Printf("A Man can be described as:----%v, With the gender as %v\n", man.description, man.gender)
}
