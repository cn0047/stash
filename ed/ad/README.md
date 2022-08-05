Algorithms & Data structures
-

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
