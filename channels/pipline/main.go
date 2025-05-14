package main

import "fmt"

func pipe[T any](inputCh <-chan T, transformF func(T) T) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)

		for v := range inputCh {
			outputCh <- transformF(v)
		}
	}()

	return outputCh
}

func main() {
	ch := make(chan string)

	transformF1 := func(x string) string {
		return fmt.Sprintf("F1 - %s", x)
	}
	transformF2 := func(x string) string {
		return fmt.Sprintf("F2 - %s", x)
	}

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%d", i)
		}
	}()

	for v := range pipe(pipe(ch, transformF1), transformF2) {
		fmt.Println(v)
	}
}
