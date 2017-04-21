package main

var (
	letters = map[uint64]string{
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thirty",
		40:   "fourty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "one hundred",
		200:  "two hundred",
		300:  "three hundred",
		400:  "four hundred",
		500:  "five hundred",
		600:  "six hundred",
		700:  "seven hundred",
		800:  "eight hundred",
		900:  "nine hundred",
		1000: "one thousand",
	}
)

func howManyLetters(n int) int {
	nWords := 0
	for i := uint64(1); i <= uint64(n); i++ {
		x, exists := letters[i]
		if exists {
			nWords += len(x)
			continue
		}
		if i < 100 {
			remaining := i % 10
			nWords += len(letters[i-remaining])
			nWords += len(letters[remaining])
		}
		if i > 100 {
			remaining := i % 100
			nWords += len(letters[i-remaining])
			nWords += len("and")
			finalRemaining := remaining % 10
			nWords += len(letters[remaining-finalRemaining])
			nWords += len(letters[finalRemaining])
		}
	}

	return nWords
}

var digits = map[int]int{
	1: 3, //"one",
	2: 3, //"two",
	3: 5, //"three",
	4: 4, //"four",
	5: 4, //"five",
	6: 3, //"six",
	7: 5, //"seven",
	8: 5, //"eight",
	9: 4, //"nine",
}

var tens = map[int]int{
	10: 3, //"ten",
	11: 6, //"eleven",
	12: 6, //"twelve",
	13: 8, //"thirteen",
	14: 8, //"fourteen",
	15: 7, //"fifteen",
	16: 7, //"sixteen",
	17: 9, //"seventeen",
	18: 8, //"eighteen",
	19: 8, //"nineteen",
	2:  6, //"twenty",
	3:  6, //"thirty",
	4:  5, //"forty",
	5:  5, //"fifty",
	6:  5, //"sixty",
	7:  7, //"seventy",
	8:  6, //"eighty",
	9:  6, //"ninety",
}

func numLetters(n int) int {
	// catch special cases like single digits or teens
	if digits[n] > 0 {
		return digits[n]
	}
	if tens[n] > 0 {
		return tens[n]
	}
	if n == 1000 {
		return 3 + 8
	}
	// special case "one hundred", "two hundred", etc
	if n%100 == 0 {
		return digits[n/100] + 7
	}
	if n > 100 {
		return digits[n/100] + 10 + numLetters(n%100)
	}
	if n%10 == 0 {
		return tens[n/10]
	}
	return tens[n/10] + digits[n%10]
}

func verifySol() uint64 {
	sum := uint64(0)
	for i := 1; i <= 1000; i++ {
		// println(i, num_letters(i))
		sum += uint64(numLetters(i))
	}
	return sum
}
