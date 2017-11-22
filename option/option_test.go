//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004

// Copyright (C) realityone <realityone@me.com> | https://reality0ne.com

// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.

//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

//  0. You just DO WHAT THE FUCK YOU WANT TO.

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
