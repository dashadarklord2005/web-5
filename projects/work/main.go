package main

import (
	"fmt"
	"sync"
)

func work() {
	
	fmt.Println("WORK!!")
}

func main() {
	var wg sync.WaitGroup 

	wg.Add(10) 

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() 
			work()          
		}()
	}

	wg.Wait() 
	fmt.Println("Ура! Все горутины завершились!")
}
