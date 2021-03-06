package main
//Defines the package for export. Default is main

//Go Comments start with two slashes
/* This is a multiline Go Comment */

//Import statements
import "fmt"
import "math/rand"

Import multiple at once
import (
  "fmt"        
  "math/rand" 
)

//Import as alias r
import r "math/rand"


// function matches package name defined at top of file
func main() {
    fmt.Println("Hello, world!")
    fmt.Println( rand.Intn(100))
}


func main() {
    fmt.Println( r.Intn(100))
}

//functions
func template() {
    fmt.Println("This a sample function");
}

Variable assignment operator is :=
Assignment 
str := "Hello"
str := `Multiline string`
num := 3                  // int
num := 3.                 // float64
num := 3 + 4i             // complex128
num := byte('a')          // byte (alias for uint8)

//Conditional
if day == "sunday" || day == "saturday" {
  rest()
} else if day == "monday" && isTired() {
  groan()
} else {
  work()
}

//For loop
for count := 0; count <= 10; count++ {
  fmt.Println("My counter is at", count)
}

//For-Range loop
entry := []string{"Jack","John","Jones"}
for i, val := range entry {
  fmt.Printf("At position %d, the character %s is present\n", i, val)
}

//while loop
n := 0
x := 42
for n != x {
  n := guess()
}

Decrement variable
n--
Increment 
n++

//Add a function passScores() that will generate 10 random test scores(0-100) and print all passing scores(>65).
//Add the call invocation of passScores() to the autogenerated main() function;
//Comment out all prior examples to run.
func passScores() {

}


//Arrays - defined with fixed size
// arr := []int{3, 7, 2, 9}

//Find the size of the array
// len(arr) == 4

//Access value at index
  // arr[2]   // 2




//Define a function called findPair() that takes an array of integers, the length of the array, and target.
//Find if any two unique elements sum is equal to the target

//Uncomment this and comment out the prior main()
// func main() {
//   arr := []int{8, 7, 2, 5, 3, 1}
// 	// arr := []int{3, 7, 2, 9}
//   findPair(arr, len(arr), 10)
// }


// func findPair(arr []int, n, sum int) {

// }


