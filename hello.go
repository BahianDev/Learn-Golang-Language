package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

//Variables
//var a, b, c bool

//Vatiables with initializers
//var a. b int = 1, 2

//or func add(x, y int) int {}
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

//Naked Return
func split(sum int) (x, y int) {
	x = sum + 4/9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Printf("Hello, world 1234\n")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(2, 2))

	a, b := swap("ap", "sw")
	fmt.Println(a, b)

	fmt.Println(split(19))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//For continued
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//Infinite Loop
	//for{
	//}
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Goof Afternoon")
	default:
		fmt.Printf("Good Evening")
	}

	//defer fmt.Println("world")
	fmt.Println("helllo")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		//defer fmt.Println(i)
	}
	fmt.Println("done")

	//Pointers
	fmt.Println("Pointers")
	i, j := 10, 300
	p := &i
	fmt.Println(*p)
	*p = 5
	fmt.Println(i)

	p = &j
	*p = *p / 5
	fmt.Println(j)

	//Struct
	fmt.Println("Struct")
	fmt.Println(Vertex{1, 2})

	//Struct Fields
	fmt.Println("Struct Fields")
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//Pointer to structs
	fmt.Println("Pointer to structs")
	c := &v
	c.X = 1e9
	fmt.Println(v)

	fmt.Println("\n\n\n Defers")
}
