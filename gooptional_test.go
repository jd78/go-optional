package gooptional

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	obj := "test"
	opt := Some(obj)

	assert.True(t, opt.IsSome())
	assert.False(t, opt.IsNone())
	assert.Equal(t, obj, opt.Get())
}

func TestNone(t *testing.T) {
	opt := None[int]()
	assert.Panics(t, func() { opt.Get() })
	assert.Panics(t, func() { opt.GetError() })
	assert.True(t, opt.IsNone())
	assert.False(t, opt.IsSome())
}

func TestError(t *testing.T) {
	err := errors.New("sample error")
	opt := Error[int](err)

	assert.True(t, opt.IsNone())
	assert.False(t, opt.IsSome())
	assert.True(t, opt.HasError())
	assert.Equal(t, err.Error(), opt.GetError().Error())
}

func TestWorksWithZeroValues(t *testing.T) {
	opt := Some[int](0)
	assert.True(t, opt.IsSome())
	assert.Equal(t, 0, opt.Get())
}

func TestPanicIfNil(t *testing.T) {
	var i *int
	assert.Panics(t, func() { Some[*int](i) })
}
