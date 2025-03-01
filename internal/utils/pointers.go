package utils

// ToPtr принимает значение и возвращает указатель на него
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr принимает указатель и возвращает значение (или zero-value, если nil)
func FromPtr[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
