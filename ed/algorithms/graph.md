Graph
-

Graph - it's mathematical structure used to model pairwise relations between objects.
Graph - it's just dots and lines connecting them.

Vertices - points.
Edges - lines.
Cardinality - count of vertices.
Degree (or valency) of vertex - number of edges that connect to it.

Multigraph - point connected with itself (has loopback).
Directed graph - all edges are directed (have direction) from one vertex to another.
Weighted graph - each edge have numerical weight. 
Complete graph - every vertex is connected to every other vertex.

All edges can be stored in adjacent matrix or adjacent list or in icidence matrix.

## Path

Path - set of edges that conncect 2 nodes.
Cycle -
Simple path - only use every edge at most once.

Hamiltonion path - path that visits each node exactly once.

Euler circuit - path has same initial and terminal veritces.
Also all vertex have even degree.

Euler path - includes exactly once all the edges and has different first and last veritces.
Also all vertex have even degree but 2 have odd degree.

Held Karp Algorithm.

## Shortest path

* Dijkstra's Algorithm (only for positive weights).

* Bellman-Ford Algorithm.

* BFS.

* Topological sort.

* Floyd-Warshall Algorithm.

## Search

BFS (Breadth First Search).
Uses queue: add node to queu -> add non visited neightors to queue -> deque top of queue.
BFS generates a spanning tree.

DFS (Depth First Search).
Uses recursion or
Uses stack: add node to stack -> ...

## Tree

To make Minimum spanning tree use:

* Prim's Algorithm - like BFS but pick chipest node on each iteration.

* Kruskal's Algorithm - connect to endes while them in different separated trees,
  because cycles not allowed.

* Boruvka's Algorithm

* Reverse-Delete Algorithm

## Topological Sorting

Topological Sorting not have to be unique.
Won't work for Directed Acyclic Graph (DAG).

* Kahn's Algorithm.

* Tarjan's Algorithm
