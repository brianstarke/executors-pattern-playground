package internal

import (
	"context"
	"log"
)

// Inmplement Executable
func (c *CreateFoo) Do(ctx context.Context) error {
	log.Printf("creating foo %s\n", c.Bar)
	log.Printf("with context %#+v", ctx)

	c.Result = &Foo{
		ID:        "mock-001",
		Bar:       c.Bar,
		Baz:       c.Baz,
		CreatedBy: ctx.Value(CKUserID).(string),
	}

	return nil
}

type CreateFoo struct {
	JWTAuthorized

	Bar string
	Baz string

	Result *Foo
}

type Foo struct {
	ID  string
	Bar string
	Baz string

	CreatedBy string
}
