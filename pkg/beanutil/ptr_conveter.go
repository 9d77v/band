package beanutil

func ToValue[T any](s *T) T {
	if s != nil {
		return *s
	}
	var t T
	return t
}

func ToPointer[T any](s T) *T {
	return &s
}
