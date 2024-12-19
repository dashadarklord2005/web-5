package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream) 

	var prevValue string 
	for value := range inputStream {
		if value != prevValue { 
			outputStream <- value 
			prevValue = value     
		}
	}
}

func main() {
	in := make(chan string)
	out := make(chan string)

	go func() {
		defer close(in) 
		in <- "one"
		in <- "two"
		in <- "two"
		in <- "three"
		in <- "three"
		in <- "three"
		in <- "four"
		in <- "five"
	}()

	go removeDuplicates(in, out)

	for value := range out {
		fmt.Println(value)
	}
}
