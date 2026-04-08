package utils

import "iter"

func Map[T any, V any](input []T, fn func(T) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, item := range input {
			if !yield(fn(item)) {
				return
			}
		}
	}
}
