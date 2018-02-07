package fn

type Any = interface{}

type Compatator interface {
	Compare(a, b Any) bool
}

type CompareFunc func(a, b Any) bool

func (cf CompareFunc) Compare(a, b Any) bool {
	return cf(a, b)
}

type Predicator interface {
	Predicate(a Any) bool
}

type PredicateFunc func(a Any) bool

func (pf PredicateFunc) Predicate(a Any) bool {
	return pf(a)
}

type BinaryPredicator interface {
	Predicate(a, b Any) bool
}

type BinaryPredicateFunc func(a, b Any) bool

func (pf BinaryPredicateFunc) Predicate(a, b Any) bool {
	return pf(a, b)
}
