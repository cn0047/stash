Sharding
-

Redis cluster helps to:
* Automatically split dataset among multiple nodes.
* Continue operations when subset of nodes are unable.

Minimal cluster that works as expected must contain at least 3 master nodes.
For deployment, strongly recommend 6 node cluster, with 3 masters and 3 replicas.

Adding new node - process of adding empty node and then moving some data into it.
Resharding - move hash slots from set of nodes to another set of nodes.
It's possible to remove replica node, but to remove master node - it must be empty.
