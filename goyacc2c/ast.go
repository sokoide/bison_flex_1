package main

import "fmt"

type Node struct {
	Kind string // "assign", "put", "if", "int", "ident" and operators
	S    string
	I    int
	N1   *Node
	N2   *Node
	Ns   []*Node
}

func newNode(kind string, s string, i int, n1, n2 *Node, ns []*Node) *Node {
	return &Node{Kind: kind, S: s, I: i, N1: n1, N2: n2, Ns: ns}
}

func evaluate(lexer *lexer) {
	evaluateNodes(lexer.program)
}

func evaluateNodes(nodes []*Node) {
	for _, n := range nodes {
		evaluateNode(n)
	}
}

func evaluateNode(node *Node) int {
	switch node.Kind {
	case "assign":
		vars[node.S] = evaluateNode(node.N1)
		return 0
	case "put":
		fmt.Println(evaluateNode(node.N1))
		return 0
	case "if":
		cond := evaluateNode(node.N1)
		if cond != 0 {
			evaluateNodes(node.Ns)
		}
		return 0
	case "int":
		return node.I
	case "ident":
		return vars[node.S]
	case "+":
		return evaluateNode(node.N1) + evaluateNode(node.N2)
	case "-":
		return evaluateNode(node.N1) - evaluateNode(node.N2)
	case "*":
		return evaluateNode(node.N1) * evaluateNode(node.N2)
	case "/":
		return evaluateNode(node.N1) / evaluateNode(node.N2)
	case "%":
		return evaluateNode(node.N1) % evaluateNode(node.N2)
	default:
		panic(fmt.Sprintf("%+v not supported", node))
	}
}
