package merkle

import "crypto/sha256"

type Node struct {
	Left  *Node
	Right *Node
	Data  []byte // most likely hash of the two txs under this node
}

func NewNode(left, right *Node, data []byte) *Node {
	node := &Node{}

	// if node is at bottom of tree
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return node
}
