package main

func list_2_12() {
	var fullName1 FullName
	fullName1.SetFullName("masanobu", "naruse")

	var fullName2 FullName = fullName1
	fullName2.SetFullName("masanobu", "sato")

	// 値オブジェクトの変更は代入操作による交換で行う
	fullName1 = fullName2

	println(fullName1.firstName, fullName1.lastName)
	println(fullName2.firstName, fullName2.lastName)
}
