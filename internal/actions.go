package internal

import (
	"context"
	"log"
)

// CreateSomething creates nothing where once there was nothing
type CreateSomething struct {
	JWTAuthorized

	ID        string
	Name      string
	CreatedBy string
}

// Inmplement Executable
func (c *CreateSomething) Do(ctx context.Context) error {
	log.Printf("creating something %s\n", c.Name)
	log.Printf("with context %#+v", ctx)

	c.ID = "mock-0001"
	c.CreatedBy = ctx.Value(CKUserID).(string)

	return nil
}
