package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PathRequest struct {
	Start Point `json:"start"`
	End   Point `json:"end"`
}

type PathResponse struct {
	Path  []Point `json:"path"`
	Error string  `json:"error,omitempty"`
}

// CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func isValidPoint(p Point, gridSize int) bool {
	return p.X >= 0 && p.X < gridSize && p.Y >= 0 && p.Y < gridSize
}

func findPathDFS(grid [][]int, start, end Point, path *[]Point, visited *[][]bool) bool {
	if !isValidPoint(start, len(grid)) || (*visited)[start.X][start.Y] {
		return false
	}

	(*visited)[start.X][start.Y] = true
	*path = append(*path, start)

	if start == end {
		return true
	}

	directions := []Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}
	for _, dir := range directions {
		next := Point{X: start.X + dir.X, Y: start.Y + dir.Y}
		if findPathDFS(grid, next, end, path, visited) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func findPathHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		sendJSONResponse(w, PathResponse{Error: "Invalid request format"}, http.StatusBadRequest)
		return
	}

	gridSize := 20
	if !isValidPoint(req.Start, gridSize) || !isValidPoint(req.End, gridSize) {
		sendJSONResponse(w, PathResponse{Error: "Invalid start or end points"}, http.StatusBadRequest)
		return
	}

	grid := make([][]int, gridSize)
	visited := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
		visited[i] = make([]bool, gridSize)
	}

	path := []Point{}
	found := findPathDFS(grid, req.Start, req.End, &path, &visited)

	if !found {
		sendJSONResponse(w, PathResponse{Error: "No path found"}, http.StatusOK)
		return
	}

	sendJSONResponse(w, PathResponse{Path: path}, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, response PathResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func main() {

	port := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/find-path", enableCORS(findPathHandler))

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s...", port)
	log.Printf("Path finder endpoint available at http://localhost%s/find-path", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
