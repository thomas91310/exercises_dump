package main

import "testing"

func TestTracker(t *testing.T) {
	it := NewIntegerTracker()
	it.track(1)
	it.track(0)

	max := it.getMax()
	if max != 1 {
		t.Errorf("Expected max of 1, but it was %v", max)
	}

	min := it.getMin()
	if min != 0 {
		t.Errorf("Expected min of 0, but it was %v", min)
	}

	mean := it.getMean()
	if mean != 0.5 {
		t.Errorf("Expected min of 0.5, but it was %v", mean)
	}

	mode := it.getMode()
	if mode != 1 {
		t.Errorf("Expected mode of 1, but it was %v", mode)
	}

	it.track(3)
	it.track(1)

	max = it.getMax()
	if max != 3 {
		t.Errorf("Expected max of 3, but it was %v", max)
	}

	min = it.getMin()
	if min != 0 {
		t.Errorf("Expected min of 0, but it was %v", min)
	}

	mean = it.getMean()
	if mean != 1.25 {
		t.Errorf("Expected min of 1.25, but it was %v", mean)
	}

	mode = it.getMode()
	if mode != 1 {
		t.Errorf("Expected mode of 1, but it was %v", mode)
	}
}
