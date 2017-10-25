package traversal

type Graph struct {
	Nodes map[int]*node
}

type node struct {
	key       int
	neighbors []*node
}

func CreateGraph() *Graph{
	return &Graph{
		Nodes: make(map[int] *node),
	}
}

func (graph *Graph) AddNode(key int) *node {
	newNode := &node{
		key: key,
	}

	graph.Nodes[key] = newNode
	return newNode
}

func (graph *Graph) AddEdge(originKey, destinationKey int) bool {
	originNode, originNodeExists := graph.Nodes[originKey]
	if !originNodeExists {
		return false
	}

	destinationNode, destinationNodeExists := graph.Nodes[destinationKey]
	if !destinationNodeExists {
		return false
	}

	originNode.neighbors = append(originNode.neighbors, destinationNode)
	return true
}

func (graph *Graph) GetNeighbors(key int) ([]*node, bool) {

	node, exists := graph.Nodes[key]
	if !exists {
		return nil, false
	}
	if len(node.neighbors)==0{
		return node.neighbors, false
	}
	return node.neighbors, true
}
