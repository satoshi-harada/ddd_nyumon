package main

import "strings"

func list_2_3() {
	var fullName = "john smith"
	var tokens = strings.Split(fullName, " ")
	var lastName = tokens[0]
	println(lastName)
}
