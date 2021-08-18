package main

import (
	"context"
	"log"

	"starke.codes/epp/internal"
)

func main() {
	c := &internal.CreateSomething{
		Name: "Test Something",
	}

	ctx := context.Background()

	err := internal.Exec(ctx, c)

	if err != nil {
		panic(err)
	}

	log.Printf("result %#+v\n", c)
}
