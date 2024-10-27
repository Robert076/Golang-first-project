package main

import "fmt" // format package, used for input and output

func main() {
	fmt.Println("Hello World")
	var x string = "Hello World"
	fmt.Println(x)
	var y = 5
	fmt.Println(y)

	fmt.Println("Printing multiple stuff", y)

	f(200)
}

func f(x int) {
	fmt.Println(x)
}
