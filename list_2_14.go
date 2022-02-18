package main

import "reflect"

func list_2_14() {
	nameA := NewFullName("masanobu", "naruse")
	nameB := NewFullName("masanobu", "naruse")
	println(reflect.DeepEqual(nameA, nameB))
}
