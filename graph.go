package main

// Vertex represents a vertex in the graph with a unique ID and a map of edges.
type Vertex struct {
	ID    string             // Unique identifier for the vertex
	Edges map[string]float64 // Map of edges from this vertex, with weights (float64)
}

// AdjacencyList represents a graph using an adjacency list.
type AdjacencyList struct {
	vertices map[string]*Vertex // Map of vertices in the graph
}

// NewAdjacencyList creates and returns a new instance of an adjacency list.
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		vertices: make(map[string]*Vertex), // Initialize the vertices map
	}
}

// AddVertex adds a new vertex with the provided nodeId to the graph.
func (g *AdjacencyList) AddVertex(nodeId string) {
	// Check if the vertex already exists; if not, add it
	
	if _, exists := g.vertices[nodeId]; !exists {
		g.vertices[nodeId] = &Vertex{
			ID:    nodeId,
			Edges: make(map[string]float64), // Initialize the edges map for the vertex
		}
	}
}

// NumVertices returns the total number of vertices in the graph.
func (g *AdjacencyList) NumVertices() int {
	return len(g.vertices) // Return the count of vertices
}

// NumEdges returns the total number of edges in the graph.
func (g *AdjacencyList) NumEdges() int {
	count := 0
	for _, vertex := range g.vertices {
		count += len(vertex.Edges) // Sum the number of edges for each vertex
	}
	return count
}

// BFS (Breadth-First Search) performs a breadth-first search on the graph from a given start node.
func (g *AdjacencyList) BFS(nodeId string) map[string]int {
	distances := make(map[string]int)
	for id := range g.vertices {
		distances[id] = -1
	}

	queue := Queue{}
	queue.Enqueue(nodeId)
	distances[nodeId] = 0

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		for neighbor := range g.vertices[current].Edges {
			if distances[neighbor] == -1 {
				queue.Enqueue(neighbor)
				distances[neighbor] = distances[current] + 1
			}
		}
	}

	return distances
}

// DFS (Depth-First Search) performs a depth-first search on the graph from a given start node.
func (g *AdjacencyList) DFS(nodeId string) map[string]bool {
	visited := make(map[string]bool)
	stack := Stack{}
	stack.Push(nodeId)

	for !stack.IsEmpty() {
		node := stack.Pop()
		if !visited[node] {
			visited[node] = true
			for neighbor := range g.vertices[node].Edges {
				if !visited[neighbor] {
					stack.Push(neighbor)
				}
			}
		}
	}

	return visited
}

// AddDirectedEdge adds a directed edge between two nodes with a specified length.
func (g *AdjacencyList) AddDirectedEdge(nodeId1, nodeId2 string, length float64) {
	if _, exists := g.vertices[nodeId1]; exists {
		g.vertices[nodeId1].Edges[nodeId2] = length // Add the edge with the length
	}
}

// Predecessors returns a slice of all predecessor nodes of a given node.
func (g *AdjacencyList) Predecessors(nodeId string) []string {
	predecessors := []string{}
	for id, vertex := range g.vertices {
		if _, exists := vertex.Edges[nodeId]; exists {
			predecessors = append(predecessors, id) // Add the predecessor node
		}
	}
	return predecessors
}

// Successors returns a slice of all successor nodes of a given node.
func (g *AdjacencyList) Successors(nodeId string) []string {
	successors := []string{}
	if vertex, exists := g.vertices[nodeId]; exists {
		for id := range vertex.Edges {
			successors = append(successors, id) // Add the successor node
		}
	}
	return successors
}

// Dijkstra computes the shortest path from a start node to all other nodes in the graph.
func (g *AdjacencyList) Dijkstra(startId string) map[string]float64 {
	maxDist := 1e9
	dist := make(map[string]float64)
	for id := range g.vertices {
		dist[id] = maxDist
	}
	dist[startId] = 0

	heap := NewMinHeap()
	heap.Push(&HeapNode{vertex: startId, dist: 0})

	for !heap.IsEmpty() {
		node := heap.Pop()
		u := node.vertex

		if dist[u] < node.dist {
			continue
		}

		for v, length := range g.vertices[u].Edges {
			if newDist := dist[u] + length; newDist < dist[v] {
				dist[v] = newDist
				heap.Push(&HeapNode{vertex: v, dist: newDist})
			}
		}
	}

	return dist
}

// TopoSort performs a topological sorting of the graph.
func (g *AdjacencyList) TopoSort() map[string]int {
    visited := make(map[string]bool)
    stack := Stack{}
    order := make(map[string]int)

    var visit func(string)
    visit = func(nodeId string) {
        if visited[nodeId] {
            return
        }
        visited[nodeId] = true

        for neighbor := range g.vertices[nodeId].Edges {
            visit(neighbor)
        }
        stack.Push(nodeId) // Push the node after visiting all neighbors
    }

    for nodeId := range g.vertices {
        visit(nodeId)
    }

    position := 0
    for !stack.IsEmpty() {
        node := stack.Pop()
        order[node] = position
        position++
    }

    return order
}


// AddUndirectedEdge adds an undirected edge between two nodes.
func (g *AdjacencyList) AddUndirectedEdge(nodeId1, nodeId2 string, length float64) {
	g.AddDirectedEdge(nodeId1, nodeId2, length) // Add edge in one direction
	g.AddDirectedEdge(nodeId2, nodeId1, length) // Add edge in the other direction
}

// Neighbors returns a slice of all neighbors of a given node.
func (g *AdjacencyList) Neighbors(nodeId string) []string {
	return g.Successors(nodeId) // Neighbors are the same as successors for undirected edges
}

// UCC finds undirected connected components in the graph.
func (g *AdjacencyList) UCC() map[string]int {
	component := make(map[string]int)
	visited := make(map[string]bool)
	queue := Queue{}
	var c int

	for nodeId := range g.vertices {
		if !visited[nodeId] {
			queue.Enqueue(nodeId)

			for !queue.IsEmpty() {
				node := queue.Dequeue()
				if !visited[node] {
					visited[node] = true
					component[node] = c

					for neighbor := range g.vertices[node].Edges {
						if !visited[neighbor] {
							queue.Enqueue(neighbor)
						}
					}
				}
			}

			c++
		}
	}

	return component
}
