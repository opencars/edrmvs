package http

import "context"

const (
	HeaderTokenID   = "X-Auth-Id"
	HeaderTokenName = "X-Auth-Name"
	HeaderUserID    = "X-User-Id"
)

type HeaderKey int

const (
	UserIDHeaderKey HeaderKey = iota
	TokenIDHeaderKey
	TokenNameHeaderKey
)

func UserIDFromContext(ctx context.Context) string {
	return ctx.Value(UserIDHeaderKey).(string)
}

func TokenIDromContext(ctx context.Context) string {
	return ctx.Value(TokenIDHeaderKey).(string)
}

func TokenNameFromContext(ctx context.Context) string {
	return ctx.Value(TokenNameHeaderKey).(string)
}

func WithUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserIDHeaderKey, id)
}

func WithTokenID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, TokenIDHeaderKey, id)
}

func WithTokenName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, TokenNameHeaderKey, name)
}
