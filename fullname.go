package main

import "reflect"

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) *FullName {
	return &FullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (f *FullName) FirstName() string {
	return f.firstName
}

func (f *FullName) LastName() string {
	return f.lastName
}

func (this *FullName) Equals(other *FullName) bool {
	if nil == other {
		return false
	}
	// 比較対象が等値（ポインター値が同じ）ならtrue
	if this == other {
		return true
	}
	// 比較対象が等価（ポインター値が異なるが値は同じ）ならtrue
	if reflect.DeepEqual(this.firstName, other.firstName) &&
		reflect.DeepEqual(this.lastName, other.lastName) {
		return true
	} else {
		return false
	}
}
