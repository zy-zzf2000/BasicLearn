package test

import (
	"Go/BasicLearn/channel"
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	ch1 := channel.Getgenerator()
	ch2 := channel.Getgenerator()

	ch := channel.Merge(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Println("ch:" + <-ch)
	}
}
