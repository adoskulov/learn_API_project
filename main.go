package main

import "fmt"

func main() {
	fmt.Println("hello git normal")
}

func GreetingFor(name string) string {
	return fmt.Sprint("Hello, %s!", name)
}