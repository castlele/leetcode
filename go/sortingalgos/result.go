package sortingalgos

type Optional[V comparable] struct {
	isNone bool
	some   V
}

func Opt[V comparable](val V) *Optional[V] {
	return &Optional[V]{isNone: false, some: val}
}

func None[V comparable]() *Optional[V] {
	return &Optional[V]{isNone: true}
}

func (opt *Optional[V]) Equals(other *Optional[V]) bool {
	if other == nil {
		return false
	}

	if opt.isNone != other.isNone {
		return false
	}

	if opt.isNone && other.isNone {
		return true
	}

	return opt.some == other.some
}
