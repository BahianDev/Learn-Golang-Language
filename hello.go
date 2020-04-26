
package main

import (
	"fmt"
	"math/rand"

)

//Variables
//var a, b, c bool

//Vatiables with initializers 
//var a. b int = 1, 2

//or func add(x, y int) int {}
func add (x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x 
}

//Naked Return 
func split(sum int) (x, y int){
	x = sum + 4 / 9
	y = sum - x
	return
}

func main(){
	fmt.Printf("Hello, world 123\n")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(2,2))

	a, b := swap("ap", "sw")
	fmt.Println(a,b)

	fmt.Println(split(19))

	sum := 0
	for i := 0; i < 10; i++{
		sum += i
	}
	fmt.Println(sum)
	
	sum=1
	for ; sum < 100;{
		sum += sum
	}
	fmt.Println(sum)
}
