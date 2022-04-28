package functionfrequency

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
)

type Function struct {
	name   string
	amount int
}

type AllFunctions []Function

func FunctionFrequency(goCode []byte) (ansOutput []string) {
	if goCode == nil {
		return nil
	}

	fileSet := token.NewFileSet()
	inputFile, err := parser.ParseFile(fileSet, "input.go", goCode, 0)
	if err != nil {
		return nil
	}

	frequencies := inspectFileForFunctions(inputFile)

	var functions AllFunctions
	for key, value := range frequencies {
		functions = append(functions, Function{name: key, amount: value})
	}

	sort.Sort(sort.Reverse(functions))

	for i, point := range functions {
		ansOutput = append(ansOutput, point.name)
		if i == 2 {
			break
		}
	}

	return ansOutput
}

func inspectFileForFunctions(inputFile *ast.File) (frequencies map[string]int) {
	frequencies = make(map[string]int)

	ast.Inspect(inputFile, func(mainNode ast.Node) bool {
		switch nodeExpr := mainNode.(type) {
		case *ast.CallExpr:
			switch exprFunc := nodeExpr.Fun.(type) {
			case *ast.Ident:
				nameFunc := exprFunc.Name
				frequencies[nameFunc] = frequencies[nameFunc] + 1
			case *ast.SelectorExpr:
				switch nestedNodeExpr := exprFunc.X.(type) {
				case *ast.Ident:
					nameFunc := nestedNodeExpr.Name + "." + exprFunc.Sel.Name
					frequencies[nameFunc] = frequencies[nameFunc] + 1
				}
			}
		}
		return true
	})

	return frequencies
}

func (a AllFunctions) Len() int           { return len(a) }
func (a AllFunctions) Less(i, j int) bool { return a[i].amount < a[j].amount }
func (a AllFunctions) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
