package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
)

func initWebgraph(t *testing.T, webgraph AdjacencyList) {
	// for a sanity check:
	//     count on the command line the number of edges and vertices by
	// grep -E -v "^#" ~/Downloads/web-Google.txt | wc -l
	// grep -E -v "^#" ~/Downloads/web-Google.txt | sed -E 's/([[:digit:]]+)[[:space:]]+([[:digit:]]+).*/\1\n\2/' | sort | uniq | wc -l

	file, _ := os.Open("./testdata/web-Google.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	start := time.Now()
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if !strings.HasPrefix((fields[0][0:1]), "#") && len(fields) == 2 {
			webgraph.AddVertex(fields[0])
			webgraph.AddVertex(fields[1])
			webgraph.AddDirectedEdge(fields[0], fields[1], 1.)
		}
		if (i % 1000000) == 0 {
			elapsed := time.Since(start)
			t.Logf("last took %s\n", elapsed)
			t.Logf("progess: %v\n", i)
			start = time.Now()
		}
		i++
	}
	t.Logf("%v lines processed\n", i)
}

// func TestWebGoogleGraph(t *testing.T) {
// 	webgraph := NewAdjacencyList()
// 	initWebgraph(t, *webgraph)

// 	// Überprüfen der Anzahl der Knoten und Kanten
// 	expectedNumVertices := 875713 // Ersetzen Sie dies durch die tatsächliche Anzahl
// 	expectedNumEdges := 5105039   // Ersetzen Sie dies durch die tatsächliche Anzahl

// 	if webgraph.NumVertices() != expectedNumVertices {
// 		t.Errorf("Expected %d vertices, got %d", expectedNumVertices, webgraph.NumVertices())
// 	}
// 	if webgraph.NumEdges() != expectedNumEdges {
// 		t.Errorf("Expected %d edges, got %d", expectedNumEdges, webgraph.NumEdges())
// 	}

// }

func TestWebGoogleBFS(t *testing.T) {
	webgraph := NewAdjacencyList()
	initWebgraph(t, *webgraph)

	startNodeID := "11342"
	bfsResult := webgraph.BFS(startNodeID)

	// Überprüfen Sie, ob die Ergebnisse plausibel erscheinen
	for node, dist := range bfsResult {
		if dist < 0 && node == startNodeID {
			t.Errorf("Found negative distance for start node %s in BFS result", node)
		}
		// Sie können auch überprüfen, ob die Distanz für erreichbare Knoten plausibel ist
		// Zum Beispiel, ob die Distanz nicht zu groß ist
	}
}

func TestWebGoogleDFS(t *testing.T) {
	webgraph := NewAdjacencyList()
	initWebgraph(t, *webgraph)

	// Wählen Sie einen Knoten für den DFS-Test
	dfsResult := webgraph.DFS("11342")

	// Setzen Sie einen realistischen Schätzwert für die Anzahl der besuchten Knoten
	expectedNumVisitedNodes := 100 // Beispielwert, anpassen nach Bedarf

	// Überprüfen Sie, ob die Ergebnisse plausibel erscheinen
	if len(dfsResult) < expectedNumVisitedNodes {
		t.Errorf("Expected at least %d visited nodes, got %d", expectedNumVisitedNodes, len(dfsResult))
	}
}

func TestWebGoogleDijkstra(t *testing.T) {
	webgraph := NewAdjacencyList()
	initWebgraph(t, *webgraph)

	// Wählen Sie einen Startknoten
	dijkstraResult := webgraph.Dijkstra("11342")

	// Überprüfen Sie, ob die Ergebnisse plausibel erscheinen
	// Beispiel: Überprüfen Sie, ob alle Distanzen >= 0 sind
	for _, dist := range dijkstraResult {
		if dist < 0 {
			t.Errorf("Found negative distance in Dijkstra result")
		}
	}
}

func TestWebGoogleTopoSort(t *testing.T) {
	webgraph := NewAdjacencyList()
	initWebgraph(t, *webgraph)

	topoSortResult := webgraph.TopoSort()

	// Da es schwierig ist, die Korrektheit der topologischen Sortierung zu überprüfen,
	// könnten Sie sich darauf konzentrieren, zu überprüfen, ob die Sortierung vollständig ist.
	if len(topoSortResult) != webgraph.NumVertices() {
		t.Errorf("Topological sort did not include all vertices")
	}
}

func TestWebGoogleUCC(t *testing.T) {
	webgraph := NewAdjacencyList()
	initWebgraph(t, *webgraph)

	uccResult := webgraph.UCC()

	// Überprüfen Sie, ob die Anzahl der Komponenten plausibel erscheint
	// Beispiel: Überprüfen Sie, ob es mindestens eine Komponente gibt
	if len(uccResult) == 0 {
		t.Errorf("Expected at least one connected component")
	}
}
