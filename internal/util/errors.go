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

package util

import "github.com/ceph/ceph-csi/internal/sentinel"

var (
	// ErrKeyNotFound is returned when requested key in omap is not found
	ErrKeyNotFound = sentinel.New("Key not found")
	// ErrObjectExists is returned when named omap is already present in rados
	ErrObjectExists = sentinel.New("Object exists")
	// ErrObjectNotFound is returned when named omap is not found in rados
	ErrObjectNotFound = sentinel.New("Object not found")
	// ErrSnapNameConflict is generated when a requested CSI snap name already exists on RBD but with
	// different properties, and hence is in conflict with the passed in CSI volume name
	ErrSnapNameConflict = sentinel.New("Snapshot name conflict")
	// ErrPoolNotFound is returned when pool is not found
	ErrPoolNotFound = sentinel.New("Pool not found")
)
