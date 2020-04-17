package main

import "testing"

//test case to check age
func TestGetAge(t *testing.T) {
	student := Student{
		id:      01,
		name:    "Paul",
		age:     35,
		contact: ContactInfo{},
	}
	age := 35
	if age != student.getAge() {
		t.Errorf("Age is %d, expected %d", student.getAge(), age)
	}
}

func TestUpdateAge(t *testing.T){
	student := Student{
		id:      01,
		name:    "Paul",
		age:     35,
		contact: ContactInfo{},
	}
	ageToUpdate := 34
	newAge := student.updateAge(ageToUpdate)

	if newAge!=ageToUpdate{
		t.Errorf("Got update age %d, expected %d",newAge,ageToUpdate)
	}
}
