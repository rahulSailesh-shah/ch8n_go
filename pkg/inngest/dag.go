package inngest

import (
	"maps"

	"github.com/google/uuid"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
)

// From = SourceNodeId
// To = TargetNodeId

type DAG struct {
	nodes    []repo.Node
	edges    []repo.Connection
	adjList  map[uuid.UUID][]uuid.UUID
	inDegree map[uuid.UUID]int
	nodeMap  map[uuid.UUID]repo.Node
}

func NewDAG(nodes []repo.Node, edges []repo.Connection) *DAG {
	dag := DAG{
		nodes:    nodes,
		edges:    edges,
		adjList:  make(map[uuid.UUID][]uuid.UUID),
		inDegree: make(map[uuid.UUID]int),
		nodeMap:  make(map[uuid.UUID]repo.Node),
	}

	dag.BuildGraph()
	return &dag
}

func (dag *DAG) BuildGraph() {
	for _, node := range dag.nodes {
		dag.adjList[node.ID] = []uuid.UUID{}
		dag.inDegree[node.ID] = 0
		dag.nodeMap[node.ID] = node
	}

	for _, edge := range dag.edges {
		dag.adjList[edge.SourceNodeID] = append(dag.adjList[edge.SourceNodeID], edge.TargetNodeID)
		dag.inDegree[edge.TargetNodeID]++
	}
}

func (dag *DAG) ValidateGraph() bool {
	visitedStack := make(map[uuid.UUID]bool)
	recStack := make(map[uuid.UUID]bool)

	var hasCycle func(uuid.UUID) bool
	hasCycle = func(node uuid.UUID) bool {
		if recStack[node] {
			return true
		}
		if visitedStack[node] {
			return false
		}
		visitedStack[node] = true
		recStack[node] = true
		for _, neighbor := range dag.adjList[node] {
			if hasCycle(neighbor) {
				return true
			}
		}
		recStack[node] = false
		return false
	}

	for _, node := range dag.nodes {
		if hasCycle(node.ID) {
			return false
		}
	}

	return true
}

func (dag *DAG) GetLevelOrder() [][]repo.Node {
	inDegCopy := make(map[uuid.UUID]int)
	maps.Copy(inDegCopy, dag.inDegree)

	levels := make([][]repo.Node, 0)

	for len(inDegCopy) > 0 {
		currLevel := make([]repo.Node, 0)
		for nodeID, inDeg := range inDegCopy {
			if inDeg == 0 {
				currLevel = append(currLevel, dag.nodeMap[nodeID])
			}
		}
		levels = append(levels, currLevel)

		for _, node := range currLevel {
			delete(inDegCopy, node.ID)
			for _, neighbor := range dag.adjList[node.ID] {
				inDegCopy[neighbor]--
			}
		}
	}
	return levels
}
