// Package hof: Higher Order Functions in Golang Using Iterators and Generics
package hof

import (
	"iter"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Square[T Number](arr []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !yield(v * v) {
				return
			}
		}
	}
}

func Cube[T Number](arr []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !yield(v * v * v) {
				return
			}
		}
	}
}

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
