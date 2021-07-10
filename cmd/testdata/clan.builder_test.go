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
		"should build orc": {
			input: func() interface{} {
				return NewOrc().Name("Joe").Age(32).Title("BigHammer").Build()
			},
			expect: &Orc{Name: "Joe", Age: 32, Title: stringPtr("BigHammer")},
		},
		"should build warg": {
			input: func() interface{} {
				return NewWarg().Name("Pup").Saddled().Breed(Warg_BREED_MOUNTAIN).Build()
			},
			expect: &Warg{Name: "Pup", Saddled: true, Breed: Warg_BREED_MOUNTAIN},
		},
		"should build caragor": {
			input: func() interface{} {
				return NewCaragor().Name("Prune").Armoured().Equipment(
					NewCaragor_Equipment().Spears(12).Bows(42).Provisions(10),
				).Build()
			},
			expect: &Caragor{Name: "Prune", Armoured: true, Equipment: &Caragor_Equipment{Spears: 12, Bows: 42, Provisions: 10}},
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

func stringPtr(val string) *string {
	return &val
}
