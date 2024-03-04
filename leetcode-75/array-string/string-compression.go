package array_string

import (
	"strconv"
)

func compress(chars []byte) int {
	if chars == nil || len(chars) == 0 {
		return 0
	} else if len(chars) == 1 {
		return 1
	}

	count := 1
	fIndex := 0
	frontChar := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] == frontChar {
			count++
		} else {
			if count != 1 {
				countStr := strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					fIndex++
					chars[fIndex] = countStr[j]
				}
				chars = append(chars[:fIndex+1], chars[i:]...)
				i = fIndex + 1
			}
			frontChar = chars[i]
			fIndex = i
			count = 1
		}
		if i == len(chars)-1 {
			if count != 1 {
				countStr := strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					fIndex++
					chars[fIndex] = countStr[j]
				}
				chars = chars[:fIndex+1]
			}
		}
	}
	return len(chars)
}
