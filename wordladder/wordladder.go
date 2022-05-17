package wordladder

import "math"

type nodeTree struct {
	father int
	child  []int
	value  string
}
type Tree struct {
	node             []nodeTree
	pointerOnReqWord []int
	seekWord         string
}

func (tree *Tree) buildUpTree(pointerOnCur int, dic []string) {
	if tree.node[pointerOnCur].value == tree.seekWord {
		tree.pointerOnReqWord = append(tree.pointerOnReqWord, pointerOnCur)
		return
	}
	var allWordsWhatFit []string
	for i, value := range dic {
		if distance(value, tree.node[pointerOnCur].value) {
			allWordsWhatFit = append(allWordsWhatFit, value)
			dic[i] = ""
		}
	}
	for _, value := range allWordsWhatFit {
		tree.node = append(tree.node, nodeTree{pointerOnCur, []int{}, value})
		supportDic := make([]string, len(dic))
		copy(supportDic, dic)
		tree.buildUpTree(len(tree.node)-1, supportDic)
	}
}

func (tree *Tree) countForMinAmountOfTransform() (result int) {
	result = math.MaxInt
	for _, pointer := range tree.pointerOnReqWord {
		curPointer, counter := pointer, 0
		for {
			counter++
			if tree.node[curPointer].father == -1 {
				break
			}
			curPointer = tree.node[curPointer].father
		}
		if result > counter {
			result = counter
		}
	}
	return result
}

func WordLadder(from string, to string, dic []string) int {
	if len(from) != len(to) || len(dic) == 0 {
		return 0
	}
	var tree Tree
	tree.node = append(tree.node, nodeTree{-1, []int{}, from})
	tree.seekWord = to
	tree.buildUpTree(0, dic)
	if len(tree.pointerOnReqWord) == 0 {
		return 0
	}
	return tree.countForMinAmountOfTransform()
}

func distance(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	var amountDif int
	for i := range first {
		if first[i] != second[i] {
			amountDif++
		}
		if amountDif > 1 {
			return false
		}
	}
	if amountDif == 0 {
		return false
	}
	return true
}
