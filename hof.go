// Package hof: Higher Order Functions in Golang Using Iterators and Generics
package hof

import (
	"iter"
	"slices"
)

// Core Array Methods

func Map[E, T any](arr []E, transform func(E) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !yield(transform(v)) {
				return
			}
		}
	}
}

func Filter[E any](arr []E, filter func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range arr {
			if filter(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Reduce[E any, T any](arr []E, fn func(T, E) T, init T) T {
	acc := init
	for _, v := range arr {
		acc = fn(acc, v)
	}
	return acc
}

func ForEach[E any](arr []E, fn func(E)) {
	for _, v := range arr {
		fn(v)
	}
}

func Find[E any](arr []E, fn func(E) bool) (E, bool) {
	var out E
	for _, v := range arr {
		if fn(v) {
			return v, true
		}
	}
	return out, false
}

func Some[E any](arr []E, fn func(E) bool) bool {
	return slices.ContainsFunc(arr, fn)
}

func Every[E any](arr []E, fn func(E) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Numeric Helpers

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Square[E Number](arr []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range arr {
			if !yield(v * v) {
				return
			}
		}
	}
}

func Cube[E Number](arr []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range arr {
			if !yield(v * v * v) {
				return
			}
		}
	}
}

func Sum[E Number](arr []E) E {
	var sum E
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Average[E Number](arr []E) float64 {
	var sum E
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func Min[E Number](arr []E) E {
	switch len(arr) {
	case 0:
		var min_ E
		return min_
	case 1:
		return arr[0]
	default:
		min_ := arr[0]
		for _, v := range arr {
			if min_ > v {
				min_ = v
			}
		}
		return min_
	}
}

func Max[E Number](arr []E) E {
	max_ := arr[0]
	for _, v := range arr {
		if max_ < v {
			max_ = v
		}
	}
	return max_
}
