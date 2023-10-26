package gooptional

import (
	"reflect"
)

func isNil[T any](t T) bool {
	v := reflect.ValueOf(t)
	kind := v.Kind()
	return (kind == reflect.Ptr ||
		kind == reflect.Interface ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Chan ||
		kind == reflect.Func) &&
		v.IsNil()
}

type Optional[T any] struct {
	hasValue bool
	val      T
	err      error
}

func Some[T any](obj T) Optional[T] {
	if isNil[T](obj) {
		panic("passed a null reference to Some")
	}
	return Optional[T]{hasValue: true, val: obj}
}

func None[T any]() Optional[T] {
	return Optional[T]{hasValue: false}
}

func Error[T any](err error) Optional[T] {
	if err == nil {
		panic("err cannot be nil")
	}
	return Optional[T]{err: err}
}

func (o Optional[T]) IsNone() bool {
	return !o.hasValue
}

func (o Optional[T]) IsSome() bool {
	return o.hasValue
}

func (o Optional[T]) HasError() bool {
	return o.err != nil
}

func (o Optional[T]) Get() T {

	if !o.hasValue {
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
