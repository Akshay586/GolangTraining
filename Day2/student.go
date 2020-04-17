package main

import (
	"fmt"
	"os"
	"strconv"
)

// ContactInfo struct
type ContactInfo struct {
	phoneNumber string
	address     string
}

// Student struct
type Student struct {
	id      int
	name    string
	age     int
	contact ContactInfo
}

func (s *Student) getAge() int {
	return s.age
}

func (s Student) toString () string{
	id := strconv.Itoa(s.id)
	age := strconv.Itoa(s.age)

	fmt.Println("ID : ", id)
	fmt.Println("Age : ", age)
	fmt.Println("Name : ", s.name)
	fmt.Println("Address : ", s.contact.address)
	fmt.Println("Phone number : ",s.contact.phoneNumber)

	return "ID : "+id + ", Name : " + s.name + ", Age : " + age + ", Phone number" +
		s.contact.phoneNumber + ", Address : " + s.contact.address
}

func (s *Student) updateAge (updatedAge int) int{
	s.age=updatedAge
	return s.age
}

// Write student to File
func writeToFile(stud Student) {
	f, err := os.Create("./studentinfo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stud.toString())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(stud.name+"'s file update successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	student := Student{005, "Akshay", 27, ContactInfo{"234763490", "Baner, Pune"}}
	newAge := student.updateAge(25)
	fmt.Println(newAge)
	writeToFile(student)
}