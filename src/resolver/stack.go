package resolver

import (
	"fmt"
)

type node map[string]bool

// A stack of Maps
type stack struct {
	nodes []node

	anode node
}

func newStack() *stack {
	o := new(stack)
	return o
}

func (n *stack) isEmpty() bool {
	return n.count() == 0
}

func (n *stack) count() int {
	return len(n.nodes)
}

func (n *stack) get(index int) node {
	return n.nodes[index]
}

func (n *stack) push(anode node) {
	n.anode = anode
	// fmt.Println("stack: pushing ", n.anode)
	n.nodes = append(n.nodes, anode)
}

func (n *stack) pop() node {
	if !n.isEmpty() {
		topI := len(n.nodes) - 1 // Top element index
		n.anode = n.nodes[topI]  // Top element
		n.nodes = n.nodes[:topI] // Pop
		return n.anode
	}

	fmt.Println("stack -- no nodes to pop")

	return nil
}

func (n *stack) top() node {
	topI := len(n.nodes) - 1
	return n.nodes[topI]
}

func (n stack) String() string {
	s := "Stack:\n"
	for _, node := range n.nodes {
		s += fmt.Sprintf("%v\n", node)
	}

	return s
}
