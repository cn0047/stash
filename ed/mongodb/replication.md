Replication
-

Durable Replica Set (write concerns):

* Unacknowledged
* Acknowledged
* Journaled - wirte data to journal (default)
* Multi-member

````
db.demo.insert({x: 1}, {writeConcern: {w: 2, j: true}});
# w: 'majority'
````

[FAQ: Replication and Replica Sets](http://docs.mongodb.org/manual/faq/replica-sets/)
