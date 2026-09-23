package main

import (
	"fmt"
	"math/rand"
)
func swap (s1,s2 string)(string, string){
	return s2,s1
}
func main(){
	fmt.Println("Hello, World!")
	x := rand.Intn(10)
	a,b := swap("hello", "world")
	fmt.Println(a,b,x)
}