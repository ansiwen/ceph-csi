/*
Copyright 2019 The Ceph-CSI Authors.

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

import "fmt"

// Error is a type for comparable errors
type Error string

func (e Error) Error() string {
	return string(e)
}

// Wrap an error in a sentinel error
func (e Error) Wrap(orig error) error {
	return errorWrap{sentinel: e, err: orig}
}

type errorWrap struct {
	sentinel Error
	err      error
}

func (e errorWrap) Error() string {
	return fmt.Sprintf("%v: %v", e.sentinel, e.err)
}

// Unwrap returns the encapsulated error.
func (e errorWrap) Unwrap() error {
	return e.err
}

// Unwrap returns the encapsulated error.
func (e errorWrap) Is(target error) bool {
	return e.sentinel == target
}
