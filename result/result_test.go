package result

import (
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	v := 123
	res := Ok(v)
	t.Logf("Result value created: %+v", res)

	if sv := res.Unwrap(); sv != v {
		t.Errorf("The stored value is not equal: %v != %v", sv, v)
	}

	or := uint32(456)
	def := uint32(789)
	err1 := fmt.Errorf("error1")
	err2 := fmt.Errorf("error2")
	if sv := Err(err1).UnwrapOr(or); sv != or {
		t.Errorf("The value is not equal in `UnwrapOr`: %v != %v", sv, or)
	}

	if sv := Err(err1).UnwrapOrElse(
		func(e error) interface{} {
			return or
		},
	); sv != or {
		t.Errorf("The value is not equal `UnwrapOrElse`: %v != %v", sv, or)
	}

	if sv := Ok(or).Map(
		func(v interface{}) interface{} {
			return v.(uint32) * 2
		},
	).Unwrap(); sv != or*2 {
		t.Errorf("The or value is not equal in `Map`: %v != %v", sv, or*2)
	}

	if sv := Ok(or).MapErr(
		func(e error) error {
			return err2
		},
	).Unwrap(); sv != or {
		t.Errorf("The or value is not equal in `MapErr`: %v != %v", sv, or)
	}

	if sv := Ok(or).And(Ok(def)).Unwrap(); sv != def {
		t.Errorf("The or value is not equal in `And`: %v != %v", sv, def)
	}

	if sv := Ok(or).AndThen(
		func(v interface{}) Result {
			return Ok(v.(uint32) * 2)
		},
	).Unwrap(); sv != or*2 {
		t.Errorf("The or value is not equal in `AndThen`: %v != %v", sv, or*2)
	}

	if sv := Ok(or).Or(Ok(def)).Unwrap(); sv != or {
		t.Errorf("The or value is not equal in `Or`: %v != %v", sv, or)
	}

	if sv := Ok(or).OrElse(
		func(e error) Result {
			return Ok(def)
		},
	).Unwrap(); sv != or {
		t.Errorf("The or value is not equal in `OrElse`: %v != %v", sv, or)
	}

	if sv := Err(err1).UnwrapOr(or); sv != or {
		t.Errorf("The or value is not equal in `UnwrapOr`: %v != %v", sv, or)
	}

	if sv := Err(err1).UnwrapOrElse(
		func(e error) interface{} {
			return def
		},
	); sv != def {
		t.Errorf("The or value is not equal in `UnwrapOrElse`: %v != %v", sv, def)
	}
}
