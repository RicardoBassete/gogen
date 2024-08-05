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
