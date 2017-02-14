// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package global

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"ui"

	"github.com/y0ssar1an/q"
)

// `StringVectorToList` converts a `ui.StringVector` to a normal `[]string`.
func StringVectorToList(lvs ui.StringVector) []string {
	sz := int(lvs.Size())
	ls := make([]string, sz)
	//var i int
	for i := 0; i < sz; i++ {
		ls = append(ls, lvs.Get(i))
	}
	return ls
}
func ListToStringVector(list []string, lvs ui.StringVector) {
	for _, s := range list {
		lvs.Add(s)
	}
}

// `SortStrings` sorts strings into order
func SortStrings(ss []string) []string {
	sort.Strings(ss)
	return ss
}

//type StringVector interface {
//Swigcptr() uintptr
//SwigIsStringVector()
//Size() (_swig_ret int64)
//Capacity() (_swig_ret int64)
//Reserve(arg2 int64)
//IsEmpty() (_swig_ret bool)
//Clear()
//Add(arg2 string)
//Get(arg2 int) (_swig_ret string)
//Set(arg2 int, arg3 string)
//}

// `EmptyStringVector`
var EmptyStringVector, CopyStringVector ui.StringVector

// `init` initialises some string vectors so they do not need to be created every time that they are needed.
func init() {
	EmptyStringVector = ui.NewStringVector()
	CopyStringVector = ui.NewStringVector()
	CopyStringVector.Add("copy")
}

// A Fanling7Error is an error-like structure that implements the ui.Error interface. This can be passed back to the C++ code
// `Fanling7Error`
type Fanling7Error struct {
	severity ui.Severity
	text     string
}

// Implement the interface.

// `Ok` is false for real errors, true for non-errors
func (e *Fanling7Error) Ok() bool { return e.severity == ui.Severity_ok }

// `Severity` is the kind of error (user, system).
func (e *Fanling7Error) Severity() ui.Severity { return e.severity }

// `Text` explains the error.
func (e *Fanling7Error) Text() string { return e.text }

var OkError ui.Error

// `init` creates `OkError` once so we do not need to do it every time.
func init() {
	OkError = ui.NewDirectorError(&Fanling7Error{severity: ui.Severity_ok, text: "ok"})
}

// `MakeError` creates an error object.
func MakeError(p interface{}) ui.Error {
	q.Q("error trapped", p)
	if ep, isFanlingError := p.(ui.Error); isFanlingError {
		log.Printf("error: %s", ep.Text)
		return ep
	}
	log.Printf("system error: %v", p)
	return ui.NewDirectorError(&Fanling7Error{severity: ui.Severity_system, text: fmt.Sprintf("%v", p)})
}

// `PanicUserError` panics with an error.
func PanicUserError(f string, p ...interface{}) {
	panic(ui.NewDirectorError(&Fanling7Error{severity: ui.Severity_user, text: fmt.Sprintf(f, p...)}))
}

// `SetupPanicHandler` sets up a panic handler (for `defer`) for a function.
func SetupPanicHandler(err *ui.Error) {
	if r := recover(); r != nil {
		*err = MakeError(r)
		TellPanic()
	}
}

// `TellPanic` outputs the stack trace for a panic.
func TellPanic() {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, false)
	log.Printf("panic is: %s\n------\n", string(buf))
}

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
