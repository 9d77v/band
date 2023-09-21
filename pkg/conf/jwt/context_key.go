package jwt

import "context"

type ContextKey struct {
	Name string
}

func NewContextKey(name string) ContextKey {
	return ContextKey{name}
}

func (c ContextKey) Get(ctx context.Context) *CustomClaims {
	claims, ok := ctx.Value(c).(*CustomClaims)
	if ok {
		return claims
	}
	return nil
}

func (c ContextKey) MustGet(ctx context.Context) *CustomClaims {
	claims, ok := ctx.Value(c).(*CustomClaims)
	if !ok {
		panic("get claims failed")
	}
	return claims
}
