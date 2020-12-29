package lib

import (
	"sync"
	"time"
)

//NodeList
type NodeList struct {
	TopNode *Node
	rwm *sync.RWMutex
}


func (nl *NodeList) Add(t time.Time) {
	now := t.Unix()

	nl.rwm.Lock()
	defer nl.rwm.Unlock()

	topNode := joinList(nl.TopNode, now)
	nl.TopNode = topNode
}


//joinList List 左边的value 总小于右边的value，然后返回最右边得到node
func joinList(n *Node, t int64) *Node {
	topNode := n
	newNode := &Node{Value: t, Size: 1}

	if n == nil {
		return newNode
	}

	if t > n.Value {
		newNode.LeftTree  =  n
		return newNode
	}

	if t == topNode.Value {
		n.Size += 1
		return topNode
	}

	for {
		if n.LeftTree == nil {
			n.LeftTree = newNode
			break
		}
		if t > n.LeftTree.Value {
			n.LeftTree, newNode.LeftTree = newNode, n.LeftTree
			break
		}
		if t == n.LeftTree.Value {
			n.LeftTree.Size += 1
			break
		}
		n = n.LeftTree
	}

	return topNode
}
