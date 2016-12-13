package multibase

import (
	"reflect"
	"fmt"
	"testing"
)

func areEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	return reflect.DeepEqual(expected, actual)
}

func toMessage(def string, msgAndArgs ...interface{}) string {
	switch {
	case msgAndArgs == nil || len(msgAndArgs) == 0:
		return def
	case len(msgAndArgs) == 1:
		return msgAndArgs[0].(string)
	default:
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if !areEqual(expected, actual) {
		t.Errorf("%s: \nexpected: %#v\nreceived: %#v\n", toMessage("Not equal", msgAndArgs...), expected, actual)
	}
}
