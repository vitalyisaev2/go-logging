// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

// defaultBackend is the backend used for all logging calls.
var defaultBackend LeveledBackend

// Backend is the interface which a log backend need to implement to be able to
// be used as a logging backend.
type Backend interface {
	Log(Level, int, *Record) error
	LogStr(Level, int, string) error
	GetFormatter() Formatter
}
