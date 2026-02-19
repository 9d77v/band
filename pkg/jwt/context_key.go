package jwt

import "context"

type ContextKey struct {
	Name string
}

func NewContextKey(name string) ContextKey {
	return ContextKey{name}
}

func (c ContextKey) Get(ctx context.Context) any {
	return ctx.Value(c)
}

func (c ContextKey) MustGet(ctx context.Context) string {
	value, ok := ctx.Value(c).(string)
	if !ok {
		panic("get claims failed")
	}
	return value
}

func (c ContextKey) GetCustomClaims(ctx context.Context) *CustomClaims {
	value, ok := ctx.Value(c).(*CustomClaims)
	if ok {
		return value
	}
	return nil
}

func (c ContextKey) MustGetCustomClaims(ctx context.Context) *CustomClaims {
	value, ok := ctx.Value(c).(*CustomClaims)
	if !ok {
		panic("get value failed")
	}
	return value
}

func (c ContextKey) GetString(ctx context.Context) *string {
	value, ok := ctx.Value(c).(*string)
	if ok {
		return value
	}
	return nil
}

func (c ContextKey) MustGetString(ctx context.Context) string {
	value, ok := ctx.Value(c).(string)
	if !ok {
		panic("get claims failed")
	}
	return value
}
