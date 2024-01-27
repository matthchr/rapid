// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rapid_test

import (
	"fmt"
	"testing"
	"time"

	"pgregory.net/rapid"
)

type S struct {
	F1   string
	Int  int
	T    time.Time
	TPtr *time.Time
}

func (s S) String() string {
	return fmt.Sprintf("%s, %d, %s", s.F1, s.Int, s.T.String())
}

func Test(t *testing.T) {
	now := time.Now()

	strGen := rapid.Just("Hello")
	intGen := rapid.IntRange(0, 100)
	timeGen := rapid.Just(now)
	sGen := rapid.MakeVariant[S](strGen.AsAny(), intGen.AsAny(), timeGen.AsAny())
	s := sGen.Example(1)

	if s.F1 != "Hello" {
		t.Errorf("Unexpected string value")
	}
	if s.Int > 100 || s.Int < 0 {
		t.Errorf("Unexpected int value")
	}
	if !s.T.Equal(now) {
		t.Errorf("Unexpected time.Time value")
	}
	if s.TPtr != nil && !s.TPtr.Equal(now) {
		t.Errorf("Unexpected time.Time ptr value")
	}
	fmt.Println(s.String())
}
