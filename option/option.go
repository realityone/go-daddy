package option

import (
	"errors"
)

var (
	ErrUnwrapNone = errors.New("option: called `Unwrap()` on a None value")
)

type DefaultFunc = func() interface{}
type MapFunc = func(interface{}) interface{}

type Option struct {
	none bool
	some interface{}
}

func New(none bool, some interface{}) Option {
	return Option{
		none: none,
		some: some,
	}
}

func Some(value interface{}) Option {
	return New(false, value)
}

func None() Option {
	return New(true, nil)
}

func (o Option) IsSome() bool {
	return !o.none
}

func (o Option) IsNone() bool {
	return o.none
}

func (o Option) Unwrap() interface{} {
	if o.none {
		panic(ErrUnwrapNone)
	}
	return o.some
}

func (o Option) UnwrapOr(or interface{}) interface{} {
	if o.none {
		return or
	}
	return o.some
}

func (o Option) UnwrapOrElse(f DefaultFunc) interface{} {
	if o.none {
		return f()
	}
	return o.some
}

func (o Option) Map(f MapFunc) Option {
	if o.none {
		return None()
	}
	return Some(f(o.some))
}

func (o Option) MapOr(u interface{}, f MapFunc) interface{} {
	if o.none {
		return u
	}
	return f(o.some)
}

func (o Option) MapOrElse(d DefaultFunc, f MapFunc) interface{} {
	if o.none {
		return d()
	}
	return f(o.some)
}
