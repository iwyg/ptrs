package ptrs

func Any[T any](v T) *T { return &v }
