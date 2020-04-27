
package main

import (
	"fmt"
	"math"
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

func sqrt(x float64) string{
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func pow (x, n, lim float64) float64{
	if v := math.Pow(x,n); v < lim{
		return v 
	}
	return lim
}
func main(){
	fmt.Printf("Hello, world 1234\n")
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

	//For continued
	sum=1 
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)

	//Infinite Loop
	//for{
	//}
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
