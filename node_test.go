package markdex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodePath(t *testing.T) {
	grandchild := newNode("grandchild", "grandchild")
	child := newParentNode("child", "child", []*Node{grandchild})
	root := newParentNode("root", "root", []*Node{child})

	assert.Equal(t, "root", root.Path())
	assert.Equal(t, "root/child", child.Path())
	assert.Equal(t, "root/child/grandchild", grandchild.Path())
}
