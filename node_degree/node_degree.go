package nodedegree

import (
	"errors"
	"fmt"
)

type FindDegree struct {
	nodes         int
	graph         [][2]int
	degreeOfGraph map[int][]int
}

func newFinder(nodes int, graph [][2]int) *FindDegree {
	return &FindDegree{
		nodes:         nodes,
		graph:         graph,
		degreeOfGraph: make(map[int][]int, 20),
	}
}

//это перестраховка на тот случай если неверные данные даются(например, {1, 2} и {2, 1})
func (finder *FindDegree) isExistAlready(what, where int) bool {
	for _, value := range finder.degreeOfGraph[where] {
		if value == what {
			return true
		}
	}
	return false
}

func (finder *FindDegree) buildData() error {
	for _, v := range finder.graph {
		if v[0] == v[1] {
			return errors.New("incorrect input data")
		}
		if !finder.isExistAlready(v[1], v[0]) {
			finder.degreeOfGraph[v[0]] = append(finder.degreeOfGraph[v[0]], v[1])
		}
		if !finder.isExistAlready(v[0], v[1]) {
			finder.degreeOfGraph[v[1]] = append(finder.degreeOfGraph[v[1]], v[0])
		}
	}
	return nil
}

func (finder *FindDegree) findAnswer(node int) (int, error) {
	if v, ok := finder.degreeOfGraph[node]; ok {
		return len(v), nil
	}
	return 0, errors.New(fmt.Sprintf("node %d not found in the graph", node))
}

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	if nodes <= 0 || len(graph) == 0 {
		return 0, errors.New("incorrect input data")
	}
	finder := newFinder(nodes, graph)
	err := finder.buildData()
	if err != nil {
		return 0, err
	}
	return finder.findAnswer(node)
}
