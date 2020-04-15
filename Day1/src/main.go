package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	starCreator()
	triangleOfNumbers()
	checkPrime(377)
	factorial(4)
}

func starCreator (){
	for i := 0; i<5; i++{
		for j := 0; j<i+1; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func triangleOfNumbers(){
	numToPrint := 1
	for i :=0;i<5; i++{
		for j:=0; j<i+1; j++{
			fmt.Print(numToPrint, " ")
			numToPrint++
		}
		fmt.Println()
	}
}

func checkPrime(check int) bool{
	for i :=2;i<=check/2;i++{
		if check%i==0 {
			fmt.Println(check, "is not a prime number")
			return false
		}
	}
	fmt.Println(check, "is a prime number")
	return true
}

func factorial(fact int) int{
	result := fact
	for i:=fact-1;i>1;i--{
		result =result*i
	}
	fmt.Println("factorial of ", fact, "is ", result)
	return fact
}