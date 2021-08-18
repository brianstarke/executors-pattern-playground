package internal

import (
	"context"
)

type Foo struct {
	ID  string
	Bar string
	Baz string

	CreatedBy string
}

type CreateFoo struct {
	JWTAuthorized

	Bar string
	Baz string

	Result *Foo
}

func (CreateFoo) TopicName() string {
	return "actions.createFoo"
}

// Inmplement Executable
func (c *CreateFoo) Do(ctx context.Context) error {
	sugar.Info("doing the work of creating foo")

	c.Result = &Foo{
		ID:        "mock-001",
		Bar:       c.Bar,
		Baz:       c.Baz,
		CreatedBy: ctx.Value(CKUserID).(string),
	}

	return nil
}
