package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func initGraph9(filename string, t *testing.T) *AdjacencyList {

	graph := NewAdjacencyList()

	file, _ := os.Open(filename) //
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",") // f[0]:id2 ,,
			var length float64         //edgeLength
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l //edgeLength{float64: l}
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])
			graph.AddDirectedEdge(id1, f[0], length)
		}

	}
	return graph
}

func TestDijkstraSmallGraph(t *testing.T) {
	// Initialize the graph with data from the test file
	graph := initGraph9("./testdata/problem9.8test.txt", t)

	// Run the Dijkstra algorithm from vertex 1
	startNodeID := "1"
	dijkstraResult := graph.Dijkstra(startNodeID)

	// Define the vertices in the graph
	totalVertices := 8 // Update this based on the number of vertices in your graph
	for i := 1; i <= totalVertices; i++ {
		vertex := strconv.Itoa(i) // Convert integer to string for the vertex ID
		if dist, exists := dijkstraResult[vertex]; exists {
			t.Logf("Shortest distance from %s to %s: %f", startNodeID, vertex, dist)
		} else {
			t.Errorf("No path from %s to %s found", startNodeID, vertex)
		}
	}
}

func TestDijkstraLargeGraph(t *testing.T) {
	// Initialize the graph with data from the large file
	graph := initGraph9("./testdata/problem9.8.txt", t)

	// Run the Dijkstra algorithm from vertex 1
	startNodeID := "1"
	dijkstraResult := graph.Dijkstra(startNodeID)

	// Vertices of interest
	verticesOfInterest := []string{"7", "37", "59", "82", "99", "115", "133", "165", "188", "197"}

	// Print the shortest-path distances for the specified vertices
	for _, vertex := range verticesOfInterest {
		if dist, exists := dijkstraResult[vertex]; exists {
			fmt.Printf("Shortest distance from %s to %s: %f\n", startNodeID, vertex, dist)
		} else {
			fmt.Printf("No path from %s to %s found\n", startNodeID, vertex)
		}
	}
}
