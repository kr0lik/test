package main

import (
	"fmt"
	"time"
)

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recover")
	//	}
	//}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recover")
			}
		}()

		panic("PANIC")
	}()

	time.Sleep(time.Second)

	fmt.Println("main")
}
