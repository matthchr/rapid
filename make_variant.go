// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rapid

import (
	"reflect"
)

// https://stackoverflow.com/questions/73864711/get-type-parameter-from-a-generic-struct-using-reflection

// MakeVariant creates a generator of values of type V, using reflection to infer the required structure.
func MakeVariant[V any](overrides ...*Generator[any]) *Generator[V] {
	var zero V
	gen := newMakeGen(reflect.TypeOf(zero), overrides)
	return newGenerator[V](&makeGen[V]{
		gen: gen,
	})
}
