package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go func() {
		escreverNaTela("Olá Mundo")
		waitGroup.Done()
	}()

	go func() {
		escreverNaTela("Programando em GO!")
		waitGroup.Done()
	}()
	waitGroup.Wait()
}

func escreverNaTela(textto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(textto)
		time.Sleep(time.Second)
	}
}
