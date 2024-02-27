package slice

import (
	"reflect"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		s         []T
		predicate func(T) (ok bool)
	}
	type testCase[T any] struct {
		name         string
		args         args[T]
		wantFiltered []T
	}
	tests := []testCase[int]{
		{
			name: "Even numbers",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				predicate: func(n int) (ok bool) {
					return n%2 == 0
				},
			},
			wantFiltered: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiltered := Filter(tt.args.s, tt.args.predicate)
			slices.Sort(gotFiltered)
			if !reflect.DeepEqual(gotFiltered, tt.wantFiltered) {
				t.Errorf("Filter() = %v, want %v", gotFiltered, tt.wantFiltered)
			}
		})
	}
}
