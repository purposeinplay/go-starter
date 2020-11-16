package context

import "context"

type contextKey string

const (
	requestIDKey = contextKey("request_id")
)

func (c contextKey) String() string {
	return "api private context key " + string(c)
}

// WithRequestID reads the request ID from the context.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// GetRequestID reads the request ID from the context.
func GetRequestID(ctx context.Context) string {
	obj := ctx.Value(requestIDKey)
	if obj == nil {
		return ""
	}

	return obj.(string)
}
