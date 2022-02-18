package main

func list_2_18() {
	nameA := NewFullName("masanobu", "naruse")
	nameB := NewFullName("john", "smith")

	println(nameA.Equals(nameB))
}
