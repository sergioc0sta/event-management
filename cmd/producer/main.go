package main

import (
	"fmt"

	"github.com/sergioc0sta/event-management/pkg/queue"
)

func main() {

	message := "First Publish"
	ch, err := queue.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	err = queue.Publish(ch, message)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message send:\n", message)

}
