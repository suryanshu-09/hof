## Higher Order Functions in Go 1.23+ using Iterators and Generics

A checklist of JavaScript-style higher-order functions implemented in Go.

(Added `Square[T Number]([]T) iter.Seq[T]` and `Cube[T Number]([]T) iter.Seq[T]`)

---

## Core Array Methods

- [x] **`Map[E, T]([]E, func(E) T) iter.Seq[T]`** — Transform each element
- [x] **`Filter[E]([]E, func(E) bool) iter.Seq[E]`** — Keep elements that satisfy a condition
- [x] **`Reduce[T, U]([]T, func(U, T) U, init U) U`** — Accumulate values into one
- [ ] **`ForEach[T]([]T, func(T))`** — Apply side-effects (printing, logging, etc.)
- [ ] **`Find[T]([]T, func(T) bool) (T, bool)`** — Return first element satisfying condition
- [ ] **`Some[T]([]T, func(T) bool) bool`** — Return `true` if any element matches
- [ ] **`Every[T]([]T, func(T) bool) bool`** — Return `true` if all elements match

---

## Numeric Helpers

- [ ] **`Sum[T Number](xs []T) T`** — Add all numbers
- [ ] **`Average[T Number](xs []T) float64`** — Compute mean
- [ ] **`Max[T Number](xs []T) T`** — Find maximum value
- [ ] **`Min[T Number](xs []T) T`** — Find minimum value

---

## Collection Utilities

- [ ] **`GroupBy[T, K comparable]([]T, func(T) K) map[K][]T`** — Cluster elements by key
- [ ] **`Partition[T]([]T, func(T) bool) ([]T, []T)`** — Split into matching/non-matching
- [ ] **`Unique[T comparable]([]T) []T`** — Remove duplicates
- [ ] **`Zip[A, B]([]A, []B) [][2]any`** — Combine two slices
- [ ] **`Unzip[A, B]([][2]any) ([]A, []B)`** — Split pairs
- [ ] **`FlatMap[T, U]([]T, func(T) []U) []U`** — Map + flatten in one step
- [ ] **`Chunk[T]([]T, size int) [][]T`** — Split slice into groups

---

## Functional Composition

- [ ] **`Compose[A, B, C](f func(B) C, g func(A) B) func(A) C`** — Compose functions (right-to-left)
- [ ] **`Pipe[A, B, C](f func(A) B, g func(B) C) func(A) C`** — Compose functions (left-to-right)
- [ ] **`Curry` patterns using closures** — Turn multi-arg func into chain of funcs

---
