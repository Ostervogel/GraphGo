package main

import (
	"testing"
)

func TestAddVertexAndEdge(t *testing.T) {
	graph := NewAdjacencyList()

	// Hinzufügen von Knoten
	vertices := []string{"A", "B", "C"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	if graph.NumVertices() != len(vertices) {
		t.Errorf("Expected %v vertices, got %v", len(vertices), graph.NumVertices())
	}

	// Hinzufügen von Kanten
	graph.AddDirectedEdge("A", "B", 1.0)
	graph.AddDirectedEdge("B", "C", 1.0)
	if graph.NumEdges() != 2 {
		t.Errorf("Expected 2 edges, got %v", graph.NumEdges())
	}
}

func TestBFS(t *testing.T) {
	graph := NewAdjacencyList()
	vertices := []string{"A", "B", "C", "D"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	graph.AddDirectedEdge("A", "B", 1)
	graph.AddDirectedEdge("B", "C", 1)
	graph.AddDirectedEdge("C", "D", 1)

	bfsResult := graph.BFS("A")
	expectedDistances := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}
	for node, dist := range expectedDistances {
		if bfsResult[node] != dist {
			t.Errorf("Expected distance from A to %v is %v, got %v", node, dist, bfsResult[node])
		}
	}
}

func TestDFS(t *testing.T) {
	graph := NewAdjacencyList()
	vertices := []string{"A", "B", "C", "D"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	graph.AddDirectedEdge("A", "B", 1)
	graph.AddDirectedEdge("B", "C", 1)
	graph.AddDirectedEdge("C", "D", 1)

	dfsResult := graph.DFS("A")
	for _, v := range vertices {
		if !dfsResult[v] {
			t.Errorf("Expected %v to be visited in DFS starting from A", v)
		}
	}
}

func TestUCC(t *testing.T) {
	graph := NewAdjacencyList()
	vertices := []string{"A", "B", "C", "D", "E"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	graph.AddUndirectedEdge("A", "B", 1)
	graph.AddUndirectedEdge("C", "D", 1)

	uccResult := graph.UCC()
	if uccResult["A"] != uccResult["B"] {
		t.Errorf("A and B should be in the same component")
	}
	if uccResult["C"] != uccResult["D"] {
		t.Errorf("C and D should be in the same component")
	}
	if uccResult["A"] == uccResult["C"] {
		t.Errorf("A and C should be in different components")
	}
}

func TestTopoSort(t *testing.T) {
	graph := NewAdjacencyList()
	vertices := []string{"A", "B", "C", "D"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	graph.AddDirectedEdge("A", "B", 1)
	graph.AddDirectedEdge("B", "C", 1)
	graph.AddDirectedEdge("C", "D", 1)

	topoSortResult := graph.TopoSort()
	if topoSortResult["A"] > topoSortResult["B"] || topoSortResult["B"] > topoSortResult["C"] || topoSortResult["C"] > topoSortResult["D"] {
		t.Errorf("Topological sort failed")
	}
}

func TestDijkstra(t *testing.T) {
	graph := NewAdjacencyList()

	// Add vertices
	vertices := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	for _, v := range vertices {
		graph.AddVertex(v)
	}

	// Add edges to the graph
	graph.AddDirectedEdge("1", "2", 2.0)
	graph.AddDirectedEdge("2", "3", 1.0)
	graph.AddDirectedEdge("1", "4", 6.0)
	graph.AddDirectedEdge("4", "3", 2.0)
	graph.AddDirectedEdge("3", "5", 1.0)
	graph.AddDirectedEdge("5", "6", 1.0)
	graph.AddDirectedEdge("6", "7", 1.0)
	graph.AddDirectedEdge("7", "8", 1.0)

	// Note: These distances are based on the above edges. Modify if the edges change.
	expectedDistances := map[string]float64{
		"1": 0, // Distance from 1 to 1
		"2": 2, // Distance from 1 to 2
		"3": 3, // Distance from 1 to 3
		"4": 6, // Distance from 1 to 4
		"5": 4, // Distance from 1 to 5
		"6": 5, // Distance from 1 to 6
		"7": 6, // Distance from 1 to 7
		"8": 7, // Distance from 1 to 8
	}

	// Run the Dijkstra algorithm from vertex "1"
	dijkstraResult := graph.Dijkstra("1")

	// Check if the distances are correct
	for vertex, expectedDist := range expectedDistances {
		if dist, exists := dijkstraResult[vertex]; exists {
			if dist != expectedDist {
				t.Errorf("Expected shortest path from 1 to %s: %f, got %f", vertex, expectedDist, dist)
			}
		} else {
			t.Errorf("Shortest path from 1 to %s not found", vertex)
		}
	}
}
