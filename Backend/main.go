package main

import (
	"encoding/json"
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
	Path []Point `json:"path"`
}

func findPathDFS(grid [][]int, start, end Point, path *[]Point, visited *[][]bool) bool {

	if start.X < 0 || start.X >= len(grid) || start.Y < 0 || start.Y >= len(grid[0]) || (*visited)[start.X][start.Y] {
		return false
	}

	(*visited)[start.X][start.Y] = true
	*path = append(*path, start)

	if start == end {
		return true
	}

	directions := []Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}
	for _, dir := range directions {
		{
			next := Point{X: start.X + dir.X, Y: start.Y + dir.Y}
			if findPathDFS(grid, next, end, path, visited) {
				return true
			}
		}

	}
	*path = (*path)[:len(*path)-1]
	return false
}

func findPathHandler(w http.ResponseWriter, r *http.Request) {
	var req PathRequest
	json.NewDecoder(r.Body).Decode(&req)

	grid := make([][]int, 20)
	for i := range grid {
		grid[i] = make([]int, 20)
	}

	visited := make([][]bool, 20)
	for i := range visited {
		visited[i] = make([]bool, 20)
	}

	path := []Point{}
	findPathDFS(grid, req.Start, req.End, &path, &visited)

	resq := PathResponse{Path: path}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resq)

}

func main() {
	http.HandleFunc("/find-path", findPathHandler)
	http.ListenAndServe(":8080", nil)
}
