package excel

type Node struct {
	OutValue string
	InValue  string
	inData   string
	outData  string
	Parent   *Node
	Children *Node
}

func (root *Node) AddNode(outValue, inValue string) {
	newNode := &Node{OutValue: outValue, InValue: inValue}
	newNode.Parent = root
	root.Children = newNode
}

func (root *Node) DeleteNode() {
	root.Parent.Children = root.Children
}
