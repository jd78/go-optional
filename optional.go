package gooptional

import "github.com/google/go-cmp/cmp"

type Optional[T any] struct {
	val T
	err error
}

func Some[T any](obj T) Optional[T] {
	var zeroT T
	if cmp.Equal(obj, zeroT) {
		panic("passed a null reference to Some")
	}
	return Optional[T]{val: obj}
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func Error[T any](err error) Optional[T] {
	if err == nil {
		panic("err cannot be nil")
	}
	return Optional[T]{err: err}
}

func (o Optional[T]) isNone() bool {
	var zeroT T
	return cmp.Equal(o.val, zeroT)
}

func (o Optional[T]) isSome() bool {
	var zeroT T
	return !cmp.Equal(o.val, zeroT)
}

func (o Optional[T]) hasError() bool {
	return o.err != nil
}

func (o Optional[T]) Get() T {
	var zeroT T
	if cmp.Equal(o.val, zeroT) {
		panic("The optional does not hold any value")
	}
	return o.val
}

func (o Optional[T]) GetError() error {
	if o.err == nil {
		panic("The optional does not hold any error")
	}
	return o.err
}
