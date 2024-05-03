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

To save graph in memory use: objects + pointers, matrix, or adjacency list.

Graph in DB:
* Adjacency list.
* Path enumeration.
* Nested sets.
* Closure table.

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

## Held Karp algorithm

...

## Shortest path

* Dijkstra's Algorithm (only for positive weights).
* Bellman-Ford Algorithm.
* BFS.
* Topological sort.
* Floyd-Warshall Algorithm.

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
  -> add Non Visited CHILD (only one child) to stack
  -> go deep and add next Non Visited CHILD to stack
    -> if no childs to add - pop element from stack
````

## Topological sorting

Topological sort - graph traversal in which each node visited after all its dependencies are visited.
Don't have to be unique.
Won't work for non Directed Acyclic Graph (DAG).

**Kahn's algorithm** - find vertices without incoming edges,
remove them from graph, repeat until: all nodes checked or there only nodes with incoming edges (cycle).
Use it to check whether graph is DAG.

* Tarjan's Algorithm.
