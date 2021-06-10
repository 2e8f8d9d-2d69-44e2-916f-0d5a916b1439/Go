package main

import (
	"fmt"

	"github.com/custom_data_types/organization"
)

func main() {
	p := organization.NewPerson("Caige", "Kelly")
	err := p.SetTwitterHandler("@name_name")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handlers: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.ID())
	println(p.FullName())
}
