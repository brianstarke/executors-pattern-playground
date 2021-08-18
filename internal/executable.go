package internal

import (
	"context"
	"errors"
	"log"
)

// Executable structs implement this interface
type Executable interface {
	Do(ctx context.Context) error
}

// Authorizeable structs check the information available in the given
// context to determine if this Executable struct can be ...executed.
type Authorizeable interface {
	CheckAuth(ctx context.Context) (bool, context.Context)
}

// JWTAuthorized is a sample implementation of Authorizeable that is also
// embeddable in to a struct.
type JWTAuthorized struct{}

func (JWTAuthorized) CheckAuth(ctx context.Context) (bool, context.Context) {
	log.Println("checking auth")

	ctx = context.WithValue(ctx, CKUserID, "USER-AWESOME")

	return true, ctx
}

func Exec(ctx context.Context, e Executable) error {
	// Chceck if we authing this
	auth, ok := e.(Authorizeable)
	if ok {
		authorized, newCtx := auth.CheckAuth(ctx)
		if !authorized {
			return errors.New("Unauthorized")
		}

		ctx = newCtx
	}

	return e.Do(ctx)
}
