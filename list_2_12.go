package main

func list_2_12() {
	fullName1 := NewFullName("masanobu", "naruse")
	fullName2 := NewFullName("masanobu", "naruse")

	// 値オブジェクトの変更は代入操作による交換で行う
	fullName1 = fullName2

	println(fullName1.firstName, fullName1.lastName)
	println(fullName2.firstName, fullName2.lastName)
}
