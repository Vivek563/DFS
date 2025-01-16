# PathFinder Project

A web application that finds a path between two points on a grid using depth-first search (DFS) algorithm. The project consists of a Go backend server and a React frontend.

## Prerequisites

-    Go (1.16 or later)
-    Node.js (14.0.0 or later)
-    npm (6.0.0 or later)

## Running the Complete Application

1. Start the backend server (in the backend directory):

```bash
go run main.go
```

2. Start the frontend development server (in the frontend directory):

```bash
npm run dev
```

3. Open your browser and navigate to `http://localhost:5173`

## How to Use

1. Click on any cell in the grid to set the start point (marked in green)
2. Click on another cell to set the end point (marked in red)
3. Click "Calculate Path" to find the path between the points (marked in blue)
4. Click "Reset Grid" to clear the grid and start over

## API Endpoints

### Find Path

-    **URL**: `/find-path`
-    **Method**: `POST`
-    **Request Body**:

```json
{
     "start": {
          "x": 0,
          "y": 0
     },
     "end": {
          "x": 5,
          "y": 5
     }
}
```

-    **Response**:

```json
{
     "path": [
          { "x": 0, "y": 0 },
          { "x": 1, "y": 0 }
     ]
}
```
