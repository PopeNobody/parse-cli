package parsecmd

import (
	"regexp"
	"strconv"
)

var numbers = regexp.MustCompile("[0-9]+")

func numericLessThan(a, b string) bool {
	aLen, bLen := len(a), len(b)
	minLen := aLen
	if bLen < minLen {
		minLen = bLen
	}
	if minLen == 0 {
		if aLen == 0 {
			return bLen != 0
		}
		return false
	}

	i := 0
	for ; i < minLen && a[i] == b[i]; i++ {
	}
	if i == minLen {
		if aLen == bLen {
			return false
		}
		if aLen == minLen {
			return true
		}
		return false
	}

	a, b = a[i:], b[i:]
	aNos, bNos := numbers.FindAllStringIndex(a, 1), numbers.FindAllStringIndex(b, 1)

	if len(aNos) == 1 && aNos[0][0] == 0 &&
		len(bNos) == 1 && bNos[0][0] == 0 {
		aNum, err := strconv.Atoi(a[:aNos[0][1]])
		if err != nil {
			return a < b
		}
		bNum, err := strconv.Atoi(b[:bNos[0][1]])
		if err != nil {
			return a < b
		}
		if aNum != bNum {
			return aNum < bNum
		}
	}
	return a < b
}

type naturalStrings []string

func (v naturalStrings) Len() int {
	return len(v)
}

func (v naturalStrings) Less(i, j int) bool {
	return !numericLessThan(v[i], v[j]) // desc order
}

func (v naturalStrings) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
