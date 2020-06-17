package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertfunc := func(t *testing.T, got int, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertfunc(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertfunc(t, got, want)

	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v ", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 3})
		want := []int{2, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v ", want, got)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 3})
		want := []int{0, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v ", want, got)
		}
	})

}
