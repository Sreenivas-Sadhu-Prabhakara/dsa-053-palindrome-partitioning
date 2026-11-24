package main

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		in   string
		want [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
		{"", [][]string{{}}},
		{"aba", [][]string{{"a", "b", "a"}, {"aba"}}},
		{"abc", [][]string{{"a", "b", "c"}}},
	}
	for _, c := range cases {
		if got := Partition(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Partition(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
