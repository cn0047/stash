Algorithms & Data structures
-

[awesome](https://github.com/tayllan/awesome-algorithms)
[awesome](https://github.com/gaerae/awesome-algorithms-education)

Algorithm solving plan:
* Clarifications about input:
  (array of integers or strings or mix, negative, duplicates, count of items in array, etc.).
* Think about input validation inside your func (empty array, blank string, etc.).
* What edge cases and pitfalls here.
* Clarifications about output.
* Think out loud about solution.
* Think about BigO (time and memory).
* Write code.
* Test code in mind using provided input.
* Run code.

Operations with sets ([1, 2, 3] [2, 3, 4]):
* union.
* intersection.
* diff.
* symmetric diff.

Factorial: `5! = 5 * 4 * 3 * 2 * 1 = 120`

Palindrome word which reads the same backward as forward: `madam, racecar, 10801`.

Non-degenerate triangle - `edge1 + edge2 > edge3`.

Binary decision diagram (BDD) -  is a data structure
that is used to represent a Boolean function.
On a more abstract level, BDDs can be considered
as a compressed representation of sets or relations.

#### [Bloom Filter](https://monosnap.com/file/pgYT6nOzrYcA9Y7Qx5Ed8EGDAO2QvX)

Is a space-efficient probabilistic data structure
used to test whether an element is a member of a set.
Query returns either "possibly in set" or "definitely not in set".
Elements can be added to the set, but not removed.
The more elements that are added to the set, the larger the probability of false positives.

#### Consistent Hashing

Consistent Hashing is a distributed hashing scheme
that operates independently of the number of servers or objects in a distributed hash table
by assigning them a position on an abstract circle (HashRing).

1. Create HashRing of all possible values.
   (You can think of the circle as all integers 0..2^32-1.)
2. Place servers on the HashRing.
3. Read/Write data:
  1. Calculate the hash `serverIndex = hashValue % n` and map it to some position on the HashRing.
  2. If server there - ok, otherwise travel clockwise on the ring until find the first server.

When a hash table is resized, only `keys/n` need to be remapped.


https://www.coursera.org/learn/algorithms-graphs-data-structures
https://www.coursera.org/learn/algorithms-divide-conquer
https://www.coursera.org/specializations/algorithms
