package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomDate interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomDate] struct {
	ID   int
	Name string
	Date T
}

type CustomMap[T comparable, V int | string] map[T]V

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	newValues := []T{}

	for _, value := range values {
		newValues = append(newValues, mapFunc(value))
	}

	return newValues
}

func main() {
	result := MapValues([]float64{1, 2, 3}, func(n float64) float64 {
		return n * 2
	})

	fmt.Println(result)

	user := User[int]{
		ID:   1,
		Name: "John",
		Date: 2025,
	}

	fmt.Println(user)

	mapped := make(CustomMap[int, string])

	mapped[1] = "one"
	mapped[2] = "two"

	dataChannel := make(chan int)

	go func() {
		dataChannel <- 100
	}()

	n := <-dataChannel

	fmt.Println(n)

	dataChannel2 := make(chan int, 1)

	dataChannel2 <- 200

	n2 := <-dataChannel2

	fmt.Println(n2)

	dataChannel2 <- 300

	n2 = <-dataChannel2

	fmt.Println(n2)

	dataChannelRange := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			dataChannelRange <- i
		}

		close(dataChannelRange)
	}()

	for it := range dataChannelRange {
		fmt.Println(it)
	}
}
