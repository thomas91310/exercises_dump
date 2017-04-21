// Write a class IntegerTracker with these methods:
// track(int) - Receives an integer for tracking.
// get_max() - Returns the max (int) of all integers seen so far.
// get_min() - Returns the min (int) of all integers seen so far.
// get_mean() - Returns the mean (float) of all integers seen so far.
// get_mode() - Returns the mode (most frequent) (int) of all integers seen so far.

//Each method should run run in constant time (O(1) time complexity).

// For example, in python:
// integer_tracker = IntegerTracker()
// integer_tracker.track(1)
// integer_tracker.track(0)
// integer_tracker.get_max() # => 1
// integer_tracker.get_min() # => 0
// integer_tracker.get_mean() # => 0.5
// integer_tracker.get_mode() # => 1 (1 or 0 is acceptable)
// integer_tracker.track(3)
// integer_tracker.track(1)
// integer_tracker.get_max() # => 3
// integer_tracker.get_min() # => 0
// integer_tracker.get_mean() # => 1.25
// integer_tracker.get_mode() # => 1
package main

import "math"

//IntegerTracker is a ...
type IntegerTracker struct {
	maxFrequency  int
	mode          int
	mean          float64
	totalTrackers int
	trackers      map[int]int
	max           int
	min           int
}

//NewIntegerTracker returns a new integer tracker object
func NewIntegerTracker() IntegerTracker {
	return IntegerTracker{
		maxFrequency:  math.MinInt64,
		mode:          math.MinInt64,
		mean:          0,
		totalTrackers: 0,
		trackers:      make(map[int]int),
		max:           math.MinInt64,
		min:           math.MaxInt64,
	}
}

func (it *IntegerTracker) track(n int) {
	frequency, exists := it.trackers[n]
	if exists {
		it.trackers[n]++
	} else {
		frequency = 1
		it.trackers[n] = frequency
	}
	if frequency > it.maxFrequency {
		it.maxFrequency = frequency
		it.mode = n
	}
	it.totalTrackers++
	it.mean = float64((float64(it.mean*float64(it.totalTrackers-1)) + float64(n)) / float64(it.totalTrackers))
	if n > it.max {
		it.max = n
	}
	if n < it.min {
		it.min = n
	}
}

func (it IntegerTracker) getMode() int {
	return it.mode
}

func (it IntegerTracker) getMax() int {
	return it.max
}

func (it IntegerTracker) getMin() int {
	return it.min
}

func (it IntegerTracker) getMean() float64 {
	return it.mean
}
