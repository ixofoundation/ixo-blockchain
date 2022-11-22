package types

type Either[L any, R any] interface {
	IsRightBiased() bool
	IsRight() bool
	Left() L
	Right() R
}

type leftImpl[L any, R any] struct {
	value L
}

func (l leftImpl[L, R]) IsRightBiased() bool { return true }
func (l leftImpl[L, R]) IsRight() bool       { return false }
func (l leftImpl[L, R]) Left() L             { return l.value }
func (l leftImpl[L, R]) Right() R            { return *new(R) }

type rightImpl[L any, R any] struct {
	value R
}

func (r rightImpl[L, R]) IsRightBiased() bool { return true }
func (r rightImpl[L, R]) IsRight() bool       { return true }
func (r rightImpl[L, R]) Left() L             { return *new(L) }
func (r rightImpl[L, R]) Right() R            { return r.value }

func FromError[R any](r R, err error) Either[error, R] {
	if err != nil {
		return Left[error, R](err)
	}
	return Right[error, R](r)
}

func Left[L any, R any](l L) Either[L, R] {
	return leftImpl[L, R]{value: l}
}

func Right[L any, R any](r R) Either[L, R] {
	return rightImpl[L, R]{value: r}
}

func RightFlatMap[L any, R any, RR any](either Either[L, R], f func(r R) Either[L, RR]) Either[L, RR] {
	if either.IsRightBiased() && either.IsRight() {
		return f(either.Right())
	}
	return Left[L, RR](either.Left())
}

func LeftFlatMap[L any, R any, LL any](either Either[L, R], f func(l L) Either[LL, R]) Either[LL, R] {
	if !either.IsRightBiased() && !either.IsRight() {
		return f(either.Left())
	}
	return Right[LL, R](either.Right())
}

// type Eiterr[L any, R any]  {

// }

// func FlatMap[L any, R any, LL func(B) A, RR any](either Either[L, R], f func(any) any) Either[LL, R] {
// 	switch {
// 		case either.IsRightBiased() && either.IsRight():
// 			return Right[LL, R](f(either.Right()))
// 	}
// 	return Either(either.Left(), either.Right())
// }

// func Map[L any, R any](either Either[L, R])(f func(R) L) {
// 	if either.IsRightBiased() && either.IsRight() {
// 		return f(either.Right())
// 	} else if either.IsRightBiased() && !either.IsRight() {
// 		return f(either.Left())
// 	}

// func Map[L any, R any](either Either[L, R])(f func(R) L) {
// 	if either.IsRightBiased() && either.IsRight() {
// 		f(either.Right())
// 	}
// }
