package main

import "fmt"

// HasCycle reports whether an undirected graph contains a cycle, using BFS:
// meeting an already-visited vertex that is not the one we arrived from
// means two distinct paths reach it.
func HasCycle(n int, edges [][2]int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	type item struct{ v, parent int }
	for s := 0; s < n; s++ {
		if visited[s] {
			continue
		}
		visited[s] = true
		queue := []item{{s, -1}}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, w := range adj[cur.v] {
				if !visited[w] {
					visited[w] = true
					queue = append(queue, item{w, cur.v})
				} else if w != cur.parent {
					return true // visited and not our parent: second path
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(HasCycle(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}))
	fmt.Println(HasCycle(3, [][2]int{{0, 1}, {1, 2}}))
}
