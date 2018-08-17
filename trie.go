package trie

// node struct
type Node struct {
	code   rune
	depth int
	children map[rune]*Node
	father  *Node
	data  []interface{}
}

// Return the father for this node.
func (n Node) Father() *Node {
	return n.father
}

// Return the data information for this node.
func (n Node) Data() []interface{} {
	return n.data
}

// Return the children for this node.
func (n Node) Children() map[rune]*Node {
	return n.children
}

// Return the code for this node
func (n Node) Code() rune {
	return n.code
}

// Return the depth for this node
func (n Node) Depth() int {
	return n.depth
}

// Create a new child for current node.
func (n *Node) newChild(code rune, data []interface{}) *Node {
	child := &Node{
		code: code,
		depth: n.depth + 1,
		children: make(map[rune]*Node),
		father: n,
		data: data,
	}
	n.children[code] = child
	return child
}

// Remove child for current node
func (n *Node) RemoveChild(code rune) {
	delete(n.children, code)
}

// prefix tree struct
type Trie struct {
	root *Node
}

// init trie
func New() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node), depth: 0},
	}
}

// Returns the root node for the Trie.
func (t *Trie) Root() *Node {
	return t.root
}

// Add node to the Trie
func (t *Trie) Add(path []rune, data []interface{}) *Node {
	node := t.root
	for i := range path {
		code := path[i]
		if n, ok := node.children[code]; ok {
			node = n
		} else {
			node = node.newChild(code, nil)
		}
	}
	node.data = data
	return node
}

// remove node to the Trie
// if this node is not the leaf node this method will delete all the children node
func (t *Trie) Remove(path []rune) {
	var i int
	node, isMatch := t.Root().findChild(path)
	if !isMatch {
		return
	}
	for f := node.Father(); f != nil; f = f.Father() {
		i++
		if len(f.Children()) > 1 {
			code := path[len(path)-i]
			f.RemoveChild(code)
		}
	}
}

func (t *Trie) Find(path []rune) (*Node, bool) {
	return t.root.findChild(path)
}

func (t *Trie) FindAllData(path []rune) ([]interface{}, bool) {
	return t.root.findChildData(path)
}

// Return path the last node for current base node
// and return the boolean for this path is in the tree
func (n *Node) findChild(path []rune) (lastNode *Node, isMatch bool) {

	if len(path) == 0 || len(n.children) == 0 {
		return n, false
	}

	baseNode := n
	for _, code := range path {
		lastNode, ok := baseNode.children[code]
		if !ok {
			return n, false
		}
		baseNode = lastNode
	}
	return baseNode, true
}

// Return path the all children data for current base node
// and return the boolean for this path is in the tree
func (n *Node) findChildData(path []rune) ([]interface{}, bool) {

	if len(path) == 0 || len(n.children) == 0 {
		return nil, false
	}

	var data []interface{}
	baseNode := n
	for _, code := range path {
		lastNode, ok := baseNode.children[code]
		if !ok {
			return data, false
		}
		data = append(data, lastNode.data)
		baseNode = lastNode
	}
	return data, true
}
