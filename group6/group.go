package group6

import (
	"github.com/micln/collection/kvk"
	"github.com/micln/collection/slice6"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Avg return the average of each slice in the map.
func Avg[K comparable, E Number](g map[K][]E) map[K]float64 {
	g2 := make(map[K]float64, len(g))
	for k, s := range g {
		g2[k] = slice6.Avg(s)
	}
	return g2
}

// Max returns the maximum of each slice in the map.
func Max[K comparable, E constraints.Ordered](g map[K][]E) map[K]E {
	g2 := make(map[K]E, len(g))
	for k, s := range g {
		g2[k] = slice6.TopBy(s, kvk.Large[E])
	}
	return g2
}

// Min returns the minimum of each slice in the map.
func Min[K comparable, E constraints.Ordered](g map[K][]E) map[K]E {
	g2 := make(map[K]E, len(g))
	for k, s := range g {
		g2[k] = slice6.TopBy(s, kvk.Less[E])
	}
	return g2
}

// TopBy returns the top elements of each slice in the map, the top is defined by the function top.
func TopBy[K comparable, E any](g map[K][]E, top func(a, b E) bool) map[K]E {
	g2 := make(map[K]E, len(g))
	for k, s := range g {
		g2[k] = slice6.TopBy(s, top)
	}
	return g2
}

// TopNBy returns the top n elements of each slice in the map, the top is defined by the function top.
func TopNBy[K comparable, E any](g map[K][]E, n int, top func(a, b E) bool) map[K][]E {
	g2 := make(map[K][]E, len(g))
	for k, s := range g {
		g2[k] = slice6.TopNBy(s, n, top)
	}
	return g2
}
