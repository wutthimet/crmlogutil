package crmlogutil

import (
	"fmt"
	"runtime"
	"strings"
)

const maxStackLength = 50

type Error struct {
	Err        error
	StackTrace string
}

type StackTrace struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`
}

func (m Error) Error() string {
	return m.Err.Error() + m.StackTrace
}

func Wrap(err error) Error {
	return Error{Err: err, StackTrace: getStackTraceError()}
}

func getStackTraceError() string {
	stackBuf := make([]uintptr, maxStackLength)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			trace = trace + fmt.Sprintf("\n\t File: %s, Line: %d. Function: %s", frame.File, frame.Line, frame.Function)
		}
		if !more {
			break
		}
	}
	return trace
}

func GetStackTrace() (stacktrace StackTrace) {
	stackBuf := make([]uintptr, maxStackLength)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	frames := runtime.CallersFrames(stack)

	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			stacktrace.Function = frame.Function
			stacktrace.File = frame.File
			stacktrace.Line = frame.Line
		}
		if !more {
			break
		}
	}
	return stacktrace
}
