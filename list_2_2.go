package main

import "strings"

func list_2_2() {
	var fullName = "naruse masanobu"
	var tokens = strings.Split(fullName, " ")
	var lastName = tokens[0]
	println(lastName)
}
