//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004

// Copyright (C) realityone <realityone@me.com> | https://reality0ne.com

// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.

//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

//  0. You just DO WHAT THE FUCK YOU WANT TO.

package result

import (
	"github.com/realityone/go-daddy/option"
)

var (
	ErrUnwrapErr = "called `Unwrap()` on an `Err` value"
)

type MapFunc = func(interface{}) interface{}
type MapErrFunc = func(error) error
type AndThenFunc = func(interface{}) Result
type OrElseFunc = func(error) Result
type UnwrapOrElseFunc = func(error) interface{}

type Result struct {
	err   error
	value interface{}
}

func New(err error, value interface{}) Result {
	return Result{
		err:   err,
		value: value,
	}
}

func Ok(value interface{}) Result {
	return New(nil, value)
}

func Err(err error) Result {
	return New(err, nil)
}

func (r Result) IsOk() bool {
	return r.err == nil
}

func (r Result) IsErr() bool {
	return r.err != nil
}

func (r Result) Ok() option.Option {
	if r.err != nil {
		return option.None()
	}
	return option.Some(r.value)
}

func (r Result) Err() option.Option {
	if r.err == nil {
		return option.None()
	}
	return option.Some(r.err)
}

func (r Result) Map(f MapFunc) Result {
	if r.err != nil {
		return Err(r.err)
	}
	return Ok(f(r.value))
}

func (r Result) MapErr(f MapErrFunc) Result {
	if r.err == nil {
		return Ok(r.value)
	}
	return Err(f(r.err))
}

func (r Result) And(res Result) Result {
	if r.err != nil {
		return Err(r.err)
	}
	return res
}

func (r Result) AndThen(f AndThenFunc) Result {
	if r.err != nil {
		return Err(r.err)
	}
	return f(r.value)
}

func (r Result) Or(res Result) Result {
	if r.err != nil {
		return res
	}
	return Ok(r.value)
}

func (r Result) OrElse(f OrElseFunc) Result {
	if r.err != nil {
		return f(r.err)
	}
	return Ok(r.value)
}

func (r Result) Unwrap() interface{} {
	if r.err != nil {
		panic(ErrUnwrapErr)
	}
	return r.value
}

func (r Result) UnwrapOr(optb interface{}) interface{} {
	if r.err != nil {
		return optb
	}
	return r.value
}

func (r Result) UnwrapOrElse(f UnwrapOrElseFunc) interface{} {
	if r.err != nil {
		return f(r.err)
	}
	return r.value
}
