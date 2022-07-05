package merkle

type Tree struct {
	RootNode *Node
}

func NewTree(txs [][]byte) *Tree {
	var nodes []Node

	// if there is an odd number of nodes, concat last two txs
	if len(txs)%2 != 0 {
		txs = append(txs, txs[len(txs)-1])
	}

	// make simple array of all txs in node format
	for _, tx := range txs {
		node := NewNode(nil, nil, tx)
		nodes = append(nodes, *node)
	}

	// go through and hash every tx, "traveling up the tree"
	// and get root hash to put into block
	for i := 0; i < len(txs)/2; i++ {
		var level []Node

		for j := 0; j < len(nodes)/2; j++ {
			node := NewNode(&nodes[j], &nodes[j+1], nil)
			level = append(level, *node)
		}

		nodes = level
	}

	// tree only contains top hash (most important!)
	tree := &Tree{&nodes[0]}

	return tree
}
