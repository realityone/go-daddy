package option

import (
	"testing"
)

func TestOptional(t *testing.T) {
	v := 123
	opt := Some(v)
	t.Logf("Optional value created: %+v", opt)

	if sv := opt.Unwrap(); sv != v {
		t.Errorf("The stored value is not equal: %v != %v", sv, v)
	}

	or := uint32(456)
	def := uint32(789)
	if sv := None().UnwrapOr(or); sv != or {
		t.Errorf("The value is not equal in `UnwrapOr`: %v != %v", sv, or)
	}

	if sv := None().UnwrapOrElse(
		func() interface{} {
			return or
		},
	); sv != or {
		t.Errorf("The value is not equal `UnwrapOrElse`: %v != %v", sv, or)
	}

	if sv := Some(or).Map(
		func(v interface{}) interface{} {
			return v.(uint32) * 2
		},
	).Unwrap(); sv != or*2 {
		t.Errorf("The or value is not equal in `Map`: %v != %v", sv, or*2)
	}

	if sv := Some(or).MapOr(def,
		func(v interface{}) interface{} {
			return v.(uint32) * 2
		},
	); sv != or*2 {
		t.Errorf("The or value is not equal in `MapOr`: %v != %v", sv, or*2)
	}

	if sv := None().MapOrElse(
		func() interface{} {
			return def
		},
		func(v interface{}) interface{} {
			return v.(uint32) * 2
		},
	); sv != def {
		t.Errorf("The or value is not equal in `MapOr`: %v != %v", sv, def)
	}
}
