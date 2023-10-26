# go-optional

Optional implementation for GoLang

The given code is an implementation of optional values in Go using generics. It defines a generic type Optional[T] which can hold an optional value of any type T.

The `Optional[T]` struct holds two fields:

The `Some[T any](obj T) Optional[T]` function creates a new Optional[T] instance with the provided non-null value obj. It first checks if obj is nil. If it is, it throws a panic indicating that a null reference was passed. Otherwise, it returns a new Optional[T] instance with val set to obj.

The `None[T any]() Optional[T]` function returns a new Optional[T] instance without a value or error. It simply returns an instance that holds no value.

The `Error[T any](err error) Optional[T]` function creates a new Optional[T] instance with the provided error value err. It checks if the err parameter is nil and throws a panic if so. Otherwise, it returns a new Optional[T] instance with err set to the provided value.

The `IsNone()` method checks if the optional value is empty.

The `IsSome()` method checks if the optional value is not empty.

The `HasError()` method checks if the optional value contains an error by checking if the err field is not nil.

The `Get()` method returns the value held by the optional. It throws a panic if the value has not been set.

The `GetError()` method returns the error held by the optional. It throws a panic if the error is not set.

## Usage

```go
opt := gooptional.Some[int](5)
opt.IsSome() //returns true
opt.IsNone() //returns false
opt.HasError() //returns false
val := opt.Get() // assign the value 5 to val
opt.GetError() //this will panic

opt2 := gooptional.None[int]()
opt2.IsSome() //returns false
opt2.IsNone() //returns true
opt2.HasError() //returns false

opt3 := gooptional.Error[int]()
opt3.IsSome() //returns false
opt3.IsNone() //returns true
opt3.HasError() //returns true
err := opt3.GetError() //assign the error to err

var *int i
opt4 := gooptional.Some(i) //this will panic
```
