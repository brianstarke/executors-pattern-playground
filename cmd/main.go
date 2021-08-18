package main

import (
	"context"
	"log"

	"starke.codes/epp/internal"
)

func main() {
	c := &internal.CreateFoo{
		Bar: "Apple",
		Baz: "Christopher",
	}

	ctx := context.Background()

	err := internal.Exec(ctx, c)

	if err != nil {
		panic(err)
	}

	log.Printf("result %#+v\n", c)
}
