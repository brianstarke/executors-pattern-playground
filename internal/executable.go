package internal

import (
	"context"
	"errors"
	"fmt"
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
	sugar.Info("checking auth")

	ctx = context.WithValue(ctx, CKUserID, "USER-AWESOME")

	return true, ctx
}

type Subscribeable interface {
	TopicName() string
}

func Exec(ctx context.Context, e Executable) error {
	// Pop an execID in the context for tracking
	ctx = context.WithValue(ctx, CKExecID, "mock-exec-id-13029381")

	// Chceck if we authing this
	auth, ok := e.(Authorizeable)
	if ok {
		authorized, newCtx := auth.CheckAuth(ctx)
		if !authorized {
			return errors.New("Unauthorized")
		}

		ctx = newCtx
	}

	sub, ok := e.(Subscribeable)
	if ok {
		sugar.Infow("publishing event",
			"topic", fmt.Sprintf("%s.%s.start", sub.TopicName(), ctx.Value(CKExecID).(string)),
			"payload", e,
		)
	}

	err := e.Do(ctx)

	if err != nil && ok {
		sugar.Infow("publishing event",
			"topic", fmt.Sprintf("%s.%s.error", sub.TopicName(), ctx.Value(CKExecID).(string)),
			"payload", e,
			"error", err,
		)
		return err
	}

	if ok {
		sugar.Infow("publishing event",
			"topic", fmt.Sprintf("%s.%s.complete", sub.TopicName(), ctx.Value(CKExecID).(string)),
			"payload", e,
		)
	}

	return err
}
