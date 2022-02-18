package main

type FullName struct {
	firstName string
	lastName  string
}

func (f *FullName) SetFullName(firstName string, lastName string) {
	f.firstName = firstName
	f.lastName = lastName
}

func list_2_4() {
	var fullName FullName
	fullName.SetFullName("masanobu", "naruse")
	println(fullName.lastName)
}
