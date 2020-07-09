/*
Copyright 2020 The Ceph-CSI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sentinel

import (
	"errors"
	"fmt"
)

// Sentinel is a type that can create comparable errors
type Sentinel interface {
	Wrap(error) error
	In(error) bool
}

type sentinel struct{ string }

// New creates a new sentinel
func New(s string) Sentinel { return &sentinel{s} }

// Wrap wraps an error in a sentinel error
func (s *sentinel) Wrap(e error) error {
	return &wrapError{sentinel: s, err: e}
}

// In method checks if a given error contains the sentinel (opposite of Is())
func (s *sentinel) In(e error) bool {
	return errors.Is(e, (*sentinelError)(s))
}

// internal type alias for being able to use sentinels as an error in errors.Is()
type sentinelError sentinel

func (e *sentinelError) Error() string {
	return string(e.string)
}

type wrapError struct {
	sentinel *sentinel
	err      error
}

func (e *wrapError) Error() string {
	return fmt.Sprintf("%v: %v", e.sentinel.string, e.err)
}

// Unwrap returns the encapsulated error.
func (e *wrapError) Unwrap() error {
	fmt.Println("Unwrap()")
	return e.err
}

// Is method for being called by errors.Is() (used internally only)
func (e *wrapError) Is(target error) bool {
	return (*sentinelError)(e.sentinel) == target
}
