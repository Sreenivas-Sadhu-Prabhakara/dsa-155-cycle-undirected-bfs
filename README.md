# Cycle in Undirected Graph (BFS)

> Striver SDE Sheet 155 · Day 23 Graph — carry each vertex's parent in the queue; any other visited neighbor is a second path, hence a cycle. Go, stdlib only, tested.

## Problem

Does an undirected graph contain a cycle? The naive "revisited = cycle" is wrong here, because every undirected edge immediately re-visits where you came from.

## Approach

BFS where each queue entry remembers *which vertex enqueued it* (its parent). Scanning a vertex's neighbors: unvisited → enqueue; visited **and not the parent** → two distinct paths reach that neighbor, which in an undirected graph is precisely a cycle. The parent exemption is the entire subtlety — it filters out the trivial back-edge along the arrival edge while still catching genuine ones. Component loop as always, so a cycle hiding in the third component is found (pinned by a test).

**Complexity:** O(V+E) time, O(V) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
