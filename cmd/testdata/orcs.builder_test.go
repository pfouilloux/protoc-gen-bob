package testdata

import (
	"reflect"
	"testing"
)

func TestBuilders(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		input  func() interface{}
		expect interface{}
	}{
		"should build location": {
			input:  func() interface{} { return NewLocation(42, 42).Build() },
			expect: &Location{Long: 42, Lat: 24},
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Helper()

			actual := tc.input()
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("mismatch\nexpected: %[1]T %[1]v\n but got: %[2]T %[2]v\n", tc.expect, actual)
			}
		})
	}
}
