package test

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, errorMsg string) {
	if a != b {
		msg := fmt.Sprintf("%v != %v", a, b)

		if len(errorMsg) == 0 {
			t.Fatal(msg)
		} else {
			t.Fatalf("%s => %s", errorMsg, msg)
		}
	}
}
