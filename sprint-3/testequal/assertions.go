//go:build !solution

package testequal

import (
	"reflect"
)

func areEqual(a, b any) bool {
	switch a.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return a == b
	case string, map[string]string, []int, []byte:
		return reflect.DeepEqual(a, b)
	default:
		return false
	}
}

func printMsg(t T, msgAndArgs ...interface{}) {
	t.Helper()
	if len(msgAndArgs) != 0 {
		msg, ok := msgAndArgs[0].(string)
		if ok {
			t.Errorf(msg, msgAndArgs[1:]...)
		} else {
			panic("Invalid message passed!")
		}
	} else {
		t.Errorf("")
	}
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true if arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if areEqual(expected, actual) {
		return true
	}

	printMsg(t, msgAndArgs...)

	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !areEqual(expected, actual) {
		return true
	}

	printMsg(t, msgAndArgs...)

	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if areEqual(expected, actual) {
		return
	}

	printMsg(t, msgAndArgs...)
	t.FailNow()
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !areEqual(expected, actual) {
		return
	}

	printMsg(t, msgAndArgs...)
	t.FailNow()
}
