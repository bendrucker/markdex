package markdex

import (
	"path/filepath"
)

// Node is an entry in a markdown content graph stored on disk
type Node struct {
	Title string
	Name  string

	Parent   *Node
	Children []*Node
}

// newNode creates a new node
func newNode(title, name string) *Node {
	return &Node{
		Title:    title,
		Name:     name,
		Children: []*Node{},
	}
}

// newParentNode creates a new node with children
func newParentNode(title, name string, children []*Node) *Node {
	node := newNode(title, name)
	for _, child := range children {
		node.AppendChild(child)
	}
	return node
}

// Root returns whether the node is a root (has no parent)
func (n *Node) Root() bool {
	return n.Parent == nil
}

// AppendChild adds a new child node and points the child to the parent
func (n *Node) AppendChild(child *Node) {
	child.Parent = n
	n.Children = append(n.Children, child)
}

// Path returns the full path to the node
func (n *Node) Path() string {
	parts := []string{}
	for ; ; n = n.Parent {
		parts = append(parts, n.Name)
		if n.Root() {
			break
		}
	}

	// reverse
	for i := len(parts)/2 - 1; i >= 0; i-- {
		opp := len(parts) - 1 - i
		parts[i], parts[opp] = parts[opp], parts[i]
	}

	return filepath.Join(parts...)
}
