package missingnumbers

import "sort"

type FinderMissing struct {
	numbers        []int
	output         []int
	gotOne, gotTwo bool
	sizeNum        int
}

func newFinderMissing(numbers []int) *FinderMissing {
	return &FinderMissing{
		numbers: numbers,
		output:  nil,
	}
}

func (finder *FinderMissing) checkArrayForMissing() {
	finder.sizeNum = len(finder.numbers)
	sort.Ints(finder.numbers)
	for i := 0; i < finder.sizeNum-1; i++ {
		if finder.numbers[i] != finder.numbers[i+1]-1 {
			finder.output = append(finder.output, finder.numbers[i+1]-1)
		}
		if finder.numbers[i] == 1 || finder.numbers[i+1] == 1 {
			finder.gotOne = true
		}
		if finder.numbers[i] == 2 || finder.numbers[i+1] == 2 {
			finder.gotTwo = true
		}
	}
}

func (finder *FinderMissing) checkOutputAndCorrect() {
	if len(finder.output) != 2 {
		if !finder.gotOne {
			finder.output = append(finder.output, 1)
			finder.gotOne = true
		}
		if !finder.gotTwo {
			finder.output = append(finder.output, 2)
			finder.gotTwo = true
		}
		if len(finder.output) == 0 {
			finder.output = []int{finder.numbers[finder.sizeNum-1] + 1, finder.numbers[finder.sizeNum-1] + 2}
		}
		if len(finder.output) == 1 {
			finder.output = append(finder.output, finder.numbers[finder.sizeNum-1]+1)
		}
	}
}

func Missing(numbers []int) (output []int) {
	if numbers == nil {
		return nil
	}
	if len(numbers) == 0 {
		return []int{1, 2}
	}

	finder := newFinderMissing(numbers)
	finder.checkArrayForMissing()
	finder.checkOutputAndCorrect()
	return finder.output
}
