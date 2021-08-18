package main

import (
	"context"

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
}
