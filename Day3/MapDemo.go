package main

import "fmt"

func main(){
	studentMap := make(map[int]string)
	studentMap[1] = "Paul"
	studentMap[2] = "Vincent"
	studentMap[3] = "Braila"
	studentMap[4] = "Jennifer"
	studentMap[5] = "Jeremy"
	studentMap[6] = "Agatha"

	for key,value := range studentMap{
		fmt.Printf("Roll number %d : %s\n",key,value)
	}

	updateStudentName(studentMap,"Jenny",4)
	fmt.Println("\nUpdated map : ",studentMap)

	removeFailures(studentMap,4)
	fmt.Println("\nAfter removing failure : ",studentMap)
}

func updateStudentName(students map[int]string, updatedName string, rollNumber int){
	for key, _ := range students {
		if key == rollNumber{
			students[key]=updatedName
		}
	}
}

func removeFailures(students map[int]string, rollNumber int){
	delete(students, rollNumber)
}
