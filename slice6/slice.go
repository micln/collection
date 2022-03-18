package slice6

import (
	"sort"

	"github.com/micln/collection/set"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

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

func FilterBy[E any](s []E, f func(E) bool) []E {
	filtered := make([]E, 0, len(s))
	for _, e := range s {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func FilterNotBy[E any](s []E, f func(E) bool) []E {
	return FilterBy(s, func(e E) bool {
		return !f(e)
	})
}

func GroupBy[E any, K comparable](s []E, f func(E) K) map[K][]E {
	m := make(map[K][]E, len(s))
	for _, e := range s {
		key := f(e)
		m[key] = append(m[key], e)
	}
	return m
}

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

func UniqBy[E any, K comparable](s []E, f func(E) K) []E {
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
