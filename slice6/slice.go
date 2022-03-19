package slice6

import (
	"sort"

	"github.com/micln/collection/kvk"
	"github.com/micln/collection/set"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Avg returns the average of the number-slice.
func Avg[E Number](s []E) (avg float64) {
	if len(s) < 1 {
		return
	}

	var sum float64
	for _, e := range s {
		sum += float64(e)
	}

	return sum / float64(len(s))
}

// Distinct remove duplicate elements from the slice.
func Distinct[E comparable](s []E) []E {
	return DistinctBy(s, kvk.Self[E])
}

// DistinctBy return a new slice with all duplicate elements removed.
func DistinctBy[E any, K comparable](s []E, f func(E) K) []E {
	return UniqueBy(s, f)
}

// FilterBy retains only those element for which the function returns true.
func FilterBy[E any](s []E, f func(E) bool) []E {
	filtered := make([]E, 0, len(s))
	for _, e := range s {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// FilterNotBy retains only those element for which the function returns false.
func FilterNotBy[E any](s []E, f func(E) bool) []E {
	return FilterBy(s, func(e E) bool {
		return !f(e)
	})
}

// GroupBy returns a map of slices, where the keys are the result of the function f applied to the elements of the slice.
func GroupBy[E any, K comparable](s []E, f func(E) K) map[K][]E {
	m := make(map[K][]E, len(s))
	for _, e := range s {
		key := f(e)
		m[key] = append(m[key], e)
	}
	return m
}

// KeyBy returns a map of slices, where the keys are the result of the function f applied to the elements of the slice.
// the duplicate elements will be removed.
func KeyBy[E any, K comparable](s []E, f func(E) K) map[K]E {
	m := make(map[K]E, len(s))
	for _, e := range s {
		key := f(e)
		m[key] = e
	}

	return m
}

// SortBy sorts the slice s in ascending order as determined by the less function.
func SortBy[E any](s []E, less func(a, b E) bool) []E {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// TopBy means sort(s).first()
func TopBy[E any](s []E, top func(a, b E) bool) (e E) {
	if len(s) < 1 {
		return
	}

	item := s[0]

	for _, e := range s {
		if top(e, item) {
			item = e
		}
	}

	return item
}

// TopNBy means sort(s).Take(n)
func TopNBy[E any](s []E, n int, top func(a, b E) bool) []E {
	if len(s) < n {
		return s
	}

	return SortBy(s, top)[:n]
}

// Unique remove duplicate elements from the slice.
func Unique[E comparable](s []E) []E {
	return UniqueBy(s, kvk.Self[E])
}

// UniqueBy returns a new slice with all duplicate elements removed.
func UniqueBy[E any, K comparable](s []E, f func(E) K) []E {
	newSlice := make([]E, 0, len(s))

	seen := set.NewHashSet[K]()
	for _, e := range s {
		key := f(e)
		if !seen.Contains(key) {
			seen.Add(key)
			newSlice = append(newSlice, e)
		}
	}

	return newSlice
}
