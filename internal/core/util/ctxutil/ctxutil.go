package ctxutil

import "context"

func FromContext[T any](ctx context.Context, key string) (T, bool) {
	val, exists := ctx.Value(key).(T)
	return val, exists
}
