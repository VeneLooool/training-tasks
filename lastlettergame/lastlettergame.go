package lastlettergame

type Finder struct {
	visited     []bool
	result      []string
	inputSlice  []string
	firstLetter map[uint8][]int
}

func newFinder(dic []string) *Finder {
	return &Finder{
		result:      nil,
		visited:     make([]bool, len(dic)),
		inputSlice:  dic,
		firstLetter: make(map[uint8][]int, 100),
	}
}

func (finder *Finder) buildData() {
	for i, value := range finder.inputSlice {
		if value != "" {
			finder.firstLetter[value[0]] = append(finder.firstLetter[value[0]], i)
		}
	}
}

func (finder *Finder) find(lastLetter uint8, curState []string) {
	if len(curState) > len(finder.result) {
		//Это костыль. Как правильно сделать? (нужно сохранить curState в result, так чтобы result не указывал на curState(так как в последующем он будет изменяться))
		finder.result = make([]string, 0)
		for _, v := range curState {
			finder.result = append(finder.result, v)
		}
	}
	for _, i := range finder.firstLetter[lastLetter] {
		if !finder.visited[i] {
			finder.visited[i] = true

			newCurState := append(curState, finder.inputSlice[i])
			finder.find(finder.inputSlice[i][len(finder.inputSlice[i])-1], newCurState)

			finder.visited[i] = false
		}
	}
}

func (finder *Finder) findLongest() {
	for i, v := range finder.inputSlice {
		finder.visited[i] = true
		finder.find(v[len(v)-1], []string{v})
		finder.visited[i] = false
	}
}
func cleanStrings(dic []string) (output []string) {
	for _, v := range dic {
		output = append(output, v[0:len(v)-1])
	}
	return output
}

func Sequence(dic []string) []string {
	if len(dic) == 0 {
		return nil
	}

	finder := newFinder(cleanStrings(dic))

	finder.buildData()
	finder.findLongest()

	return finder.result
}
