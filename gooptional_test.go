package gooptional

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	obj := "test"
	opt := Some(obj)

	assert.True(t, opt.isSome())
	assert.False(t, opt.isNone())
	assert.Equal(t, obj, opt.Get())
}

func TestNone(t *testing.T) {
	opt := None[int]()
	assert.Panics(t, func() { opt.Get() })
	assert.Panics(t, func() { opt.GetError() })
	assert.True(t, opt.isNone())
	assert.False(t, opt.isSome())
}

func TestError(t *testing.T) {
	err := errors.New("sample error")
	opt := Error[int](err)

	assert.True(t, opt.isNone())
	assert.False(t, opt.isSome())
	assert.True(t, opt.hasError())
	assert.Equal(t, err.Error(), opt.GetError().Error())
}
