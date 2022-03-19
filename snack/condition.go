package snack

import "golang.org/x/exp/constraints"

func Max[E constraints.Ordered](a, b E) E {
	return IfThen(a < b, b, a)
}

func Min[E constraints.Ordered](a, b E) E {
	return IfThen(a < b, a, b)
}

func IfThen[E any](condition bool, ifValue E, elseValue E) E {
	if condition {
		return ifValue
	}
	return elseValue
}
