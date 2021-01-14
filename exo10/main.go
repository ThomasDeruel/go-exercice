package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func maFonction(wg *sync.WaitGroup) {
	fmt.Println("J'ai fini")
	wg.Done()
}
func main(){
	// add a delta -> delta = 1
	wg.Add(1) 
	go maFonction(&wg)
	// Wait blocks until the WaitGroup counter is zero
	// ( delta must be equel to 0 )
	wg.Wait()
	fmt.Println("Fin du programme")
}