package twonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	node1a := ListNode{2, nil}
	node1b := ListNode{4, nil}
	node1c := ListNode{3, nil}

	node1a.Next = &node1b
	node1b.Next = &node1c

	node2a := ListNode{5, nil}
	node2b := ListNode{6, nil}
	node2c := ListNode{4, nil}

	node2a.Next = &node2b
	node2b.Next = &node2c

	resulta := ListNode{7, nil}
	resultb := ListNode{0, nil}
	resultc := ListNode{8, nil}

	resulta.Next = &resultb
	resultb.Next = &resultc

	resultNode := addTwoNumbers(&node1a, &node2a)

	assert.Equal(t, 7, resultNode.Val)
}
