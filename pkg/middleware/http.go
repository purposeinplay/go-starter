package middleware

import (
	"context"
	"net/http"

	wctx "github.com/oakeshq/go-starter/context"
	"github.com/pborman/uuid"
)

func RequestIDCtx(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	id := uuid.NewRandom().String()
	ctx := r.Context()
	ctx = wctx.WithRequestID(ctx, id)
	return ctx, nil
}
