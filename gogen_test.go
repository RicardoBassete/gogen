package gogen

import (
	"testing"
)

func TestIf(t *testing.T) {

	{
		int1, int2 := 1, 5
		expected, got := int1, If(true, int1, int2)
		if got != expected {
			t.Errorf("[int] Expected %d, got: %d", expected, got)
		}
		expected, got = int2, If(false, int1, int2)
		if got != expected {
			t.Errorf("[int] Expected %d, got: %d", expected, got)
		}
	}

	{
		string1, string2 := "first", "second"
		expected, got := string1, If(true, string1, string2)
		if got != expected {
			t.Errorf("[string] Expected %s, got: %s", expected, got)
		}
		expected, got = string2, If(false, string1, string2)
		if got != expected {
			t.Errorf("[string] Expected %s, got: %s", expected, got)
		}
	}

	{
		bool1, bool2 := true, false
		expected, got := bool1, If(true, bool1, bool2)
		if got != expected {
			t.Errorf("[string] Expected %v, got: %v", expected, got)
		}
		expected, got = bool2, If(false, bool1, bool2)
		if got != expected {
			t.Errorf("[string] Expected %v, got: %v", expected, got)
		}
	}
}

func TestGetLast(t *testing.T) {
	{
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected, got := 10, GetLast(intSlice)
		if got != expected {
			t.Errorf("[int] Expected %v, got: %v", expected, got)
		}

		expected, got = 0, GetLast([]int{})
		if got != expected {
			t.Errorf("[int] Expected %v, got: %v", expected, got)
		}
	}

	{
		stringSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		expected, got := "j", GetLast(stringSlice)
		if got != expected {
			t.Errorf("[string] Expected %v, got: %v", expected, got)
		}

		expected, got = "", GetLast([]string{})
		if got != expected {
			t.Errorf("[string] Expected %v, got: %v", expected, got)
		}
	}
}

func TestCompArr(t *testing.T) {
	{
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		intSlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		expected, got := true, CompArr(intSlice, intSlice2)
		if expected != got {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}

	{
		stringSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		stringSlice2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

		expected, got := true, CompArr(stringSlice, stringSlice2)
		if expected != got {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}

	{
		floatSlice := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
		floatSlice2 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

		expected, got := true, CompArr(floatSlice, floatSlice2)
		if expected != got {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}

	{
		floatSlice := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
		floatSlice2 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}

		expected, got := true, CompArr(floatSlice, floatSlice2)
		if expected == got {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}

	{
		floatSlice := []float64{}
		floatSlice2 := []float64{}

		expected, got := true, CompArr(floatSlice, floatSlice2)
		if expected != got {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}
}

func TestSplitSlice(t *testing.T) {
	{
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s1, s2 := SplitSlice(intSlice)
		if !CompArr(s1, []int{1, 2, 3, 4, 5}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
		if !CompArr(s2, []int{6, 7, 8, 9, 10}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

		s1 = append(s1, s2...)
		if !CompArr(s1, intSlice) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

		s1, s2 = SplitSlice(intSlice)
		if CompArr(s1, []int{}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
		if CompArr(s2, []int{}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

	}

	{
		stringSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

		s1, s2 := SplitSlice(stringSlice)
		if !CompArr(s1, []string{"a", "b", "c", "d", "e"}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
		if !CompArr(s2, []string{"f", "g", "h", "i", "j"}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

		s1 = append(s1, s2...)
		if !CompArr(s1, stringSlice) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

		s1, s2 = SplitSlice(stringSlice)
		if CompArr(s1, []string{}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
		if CompArr(s2, []string{}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}

	{
		floatSlice := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

		s1, s2 := SplitSlice(floatSlice)
		if !CompArr(s1, []float64{1.0, 2.0, 3.0, 4.0, 5.0}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
		if !CompArr(s2, []float64{6.0, 7.0, 8.0, 9.0, 10.0}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

		s1 = append(s1, s2...)
		if !CompArr(s1, floatSlice) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}

		s1, s2 = SplitSlice(floatSlice)
		if CompArr(s1, []float64{}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
		if CompArr(s2, []float64{}) {
			t.Errorf("[bool] Expected %v, got: %v", true, false)
		}
	}
}

func TestFilter(t *testing.T) {
	{
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expectedEvenSlice := []int{2, 4, 6, 8, 10}
		expectedOddSlice := []int{1, 3, 5, 7, 9}

		evenSlice := Filter(intSlice, func(i int) bool {
			return i%2 == 0
		})

		oddSlice := Filter(intSlice, func(i int) bool {
			return i%2 == 1
		})

		if !CompArr(evenSlice, expectedEvenSlice) {
			t.Errorf("[int] Expected %v, got: %v", expectedEvenSlice, evenSlice)
		}

		if !CompArr(oddSlice, expectedOddSlice) {
			t.Errorf("[int] Expected %v, got: %v", expectedOddSlice, oddSlice)
		}
	}
	{
		stringSlice := []string{"a", "b", "c", "a", "b", "c", "a", "b", "c"}
		expectedSliceA := []string{"a", "a", "a"}
		expectedSliceB := []string{"b", "b", "b"}
		expectedSliceC := []string{"c", "c", "c"}

		sliceA := Filter(stringSlice, func(i string) bool {
			return i == "a"
		})
		sliceB := Filter(stringSlice, func(i string) bool {
			return i == "b"
		})
		sliceC := Filter(stringSlice, func(i string) bool {
			return i == "c"
		})

		if !CompArr(sliceA, expectedSliceA) {
			t.Errorf("[string] Expected %v, got: %v", expectedSliceA, sliceA)
		}
		if !CompArr(sliceB, expectedSliceB) {
			t.Errorf("[string] Expected %v, got: %v", expectedSliceB, sliceB)
		}
		if !CompArr(sliceC, expectedSliceC) {
			t.Errorf("[string] Expected %v, got: %v", expectedSliceC, sliceC)
		}
	}
	{
		type Value struct {
			val string
		}
		structSlice := []Value{
			{val: "a"},
			{val: "b"},
			{val: "c"},
			{val: "a"},
			{val: "b"},
			{val: "c"},
			{val: "a"},
			{val: "b"},
			{val: "c"},
		}

		expectedStructSliceA := []Value{
			{val: "a"},
			{val: "a"},
			{val: "a"},
		}

		structSliceA := Filter(structSlice, func(a Value) bool {
			return a.val == "a"
		})

		for i, result := range structSliceA {
			if result != expectedStructSliceA[i] {
				t.Errorf("[struct] Expected %v, got: %v", expectedStructSliceA, structSliceA)
			}
		}

	}
}
