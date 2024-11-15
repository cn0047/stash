Graph
-

<br>Graph - mathematical structure used to model pairwise relations between objects.
<br>Graph - just dots and lines connecting them.
<br>Vertices - points.
<br>Cardinality of graph - count of vertices.
<br>Edges - lines.
<br>Degree (or valency) of vertex - number of edges that connect to it.
<br>Multigraph - point connected with itself (has loopback).
<br>Directed graph - all edges are directed (have direction) from one vertex to another.
<br>Weighted graph - each edge have numerical weight.
<br>Complete graph - every vertex is connected to every other vertex.
<br>Spanning tree - subset of undirected graph that contains all vertices of graph connected with minimum number of edges.

To save graph in memory use: objects + pointers, adjacency matrix or adjacency list.
Adjacency list space-efficient than adjacency matrix.
Adjacency matrix allows to check whether two vertices are adjacent to each other in constant time.

Graph in DB:
* Adjacency list.
* Path enumeration.
* Nested sets.
* Closure table.

Tarjan's algorithm.
Held Karp algorithm.

**Topological sort** - graph traversal in which each node visited after all its dependencies are visited.
It helps to find shortest path through weighted DAG, won't work for non DAG.
Result path don't have to be unique.

**Kahn's algorithm** - find vertices without incoming edges,
remove them from graph, repeat until: all nodes checked or there only nodes with incoming edges (cycle).
Use it to check whether graph is DAG or not.

**Prim’s algorithm** - one of the efficient methods to find the minimum spanning tree.
Number of edges = number of vertices - 1.
Algorithm: 1) select min edge, 2) select next connected min edge, 3) repeat 2 until all vertices visited.

**Kruskal’s algorithm** - .
Algorithm: 1) select min edge, 2) repeat 2, if edge creates cycle - skip it, repeat 2 until all vertices visited.

**Dijkstra's Algorithm** - find the shortest paths between nodes in a weighted graph.
Works only for positive weights.

## Shortest path

* Bellman-Ford Algorithm.
* BFS.
* Topological sort.
* Floyd-Warshall Algorithm.

## Path

<br>Path - set of edges that conncect 2 nodes.
<br>Cycle - path that starts and ends at the same vertex, without repeating any other vertices.
<br>Simple path - only use every edge at most once.
<br>Hamiltonion path - path that visits each node exactly once.
<br>Euler circuit - path has same initial and terminal veritces. Also all veritces have even degree.
<br>Euler path - includes exactly once all the edges and has different first and last veritces.
Also all vertex have even degree but 2 have odd degree.

## DAG

Directed Acyclic Graph -  finite directed graph with no directed cycles.
DAG always topologically ordered.

## Search

BFS (Breadth First Search) - `O(|V| + |E|)` (V - vertices, E - edges).
BFS generates a spanning tree.

Uses queue:
````
add node to queu
  -> add ALL non visited CHILDs to queue
  -> deque top of queue
````

DFS (Depth First Search) - `O(|V| + |E|)` (V - vertices, E - edges).

Uses recursion or stack:
````
add node to stack
  -> add non visited CHILD (only one child) to stack
  -> go deep and add next non visited CHILD to stack
  -> if no childs to add - pop element from stack
````
