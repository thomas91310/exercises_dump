package main

type sequence struct {
	timestamp int
	length    int
	remaining int
}

func sessionize(times []int, timePeriod int) []sequence {
	out := make([]sequence, len(times))
	idxPrevSequence := 0
	for idx, t := range times {
		out[idx] = sequence{
			timestamp: t,
			length:    0,
			remaining: 0,
		}
		lengthSequence := times[idx] - times[idxPrevSequence]
		if lengthSequence > timePeriod {
			fillOutput(times, out, idxPrevSequence, idx)
			idxPrevSequence = idx
		}
	}
	return out
}

func fillOutput(times []int, out []sequence, startSequence int, endSequence int) {
	lengthSequence := times[endSequence-1] - times[startSequence]
	for i := startSequence; i < endSequence; i++ {
		out[i].length = lengthSequence
		out[i].remaining = times[endSequence-1] - times[i]
	}
}
