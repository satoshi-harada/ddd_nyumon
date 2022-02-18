package main

func list_2_15() {
	nameA := NewFullName("masanobu", "naruse")
	nameB := NewFullName("john", "smith")

	compareResult := nameA.firstName == nameB.firstName &&
		nameA.lastName == nameB.lastName
	println(compareResult)
}
