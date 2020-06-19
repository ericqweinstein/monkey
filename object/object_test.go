package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello, world!"}
	hello2 := &String{Value: "Hello, world!"}

	diff1 := &String{Value: "My name is Oxnard Montalvo"}
	diff2 := &String{Value: "My name is oxnard montalvo"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Strings with identical content have different hash keys.")
	}

	if diff1.HashKey() == diff2.HashKey() {
		t.Errorf("Strings with different content have the same hash key.")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("`true` does not have the same hash as `true`, but should.")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("`false` does not have the same hash as `false`, but should.")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("`true` and `false` hash to the same value, but should not.")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	two1 := &Integer{Value: 2}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("Identical integers hash to different values, but shouldn't.")
	}

	if one1.HashKey() == two1.HashKey() {
		t.Errorf("Different integers hash to the same value, but should not.")
	}
}
