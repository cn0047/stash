Complexity of algorithms
-

Important to find upper bound for algorithm.
<br>`time` - it isn't possible to estimate time because it's vary on different machines (CPU, memory, etc.)
<br>`memory (space)`
<br>`count of steps (actions)`

Big O - Dependency between count of iterations and input parameters.

Drop constants in big O. Not O(2n), not O(3n).
Ignore the base of logs (log2, log10).
Just looking for how thing scale roughtly (lineral, quadratic, etc).

````
# time complexity
O(1)      # constant time
O(log(n)) # logarithmic time
O(n)      # linear time
O(n^2)    # n-squared time
O(2^n)    # exponential time
O(n!)     # n-factorial time

# space complexity - same like time complexity ↑

O(N^2 + N^2)            = O(N^2)
O(N^2 + N)              = O(N^2)
O(N + log N)            = O(N) # because log N < N
O(N * log N)            = O(N) # ↑
O(5 * 2^N + 10 * N^100) = O(2^N)
O(N^3 + N^2 + N)        = O(N^3)
O(N^2 + B)              = O(N^2 + B) # because we know nothing about B
O(log 2 N)              = O(log N)
O(N/2)                  = O(N)
O(N^2/2)                = O(N^2)

O(2^N)                  = for fibonacci and for recursion (with creating forks with n levels deep)
O(N!)                   = for recursion / permutations / salesman problems
O(log N)                = if divide in 1/2 or myltiply by 2
````

For array inside another array - `Big O(a * b)`,
whre a - 1st array length, b - 2nd.

If algorithm takes half of elements on iteration - `Big O = O(log N)`.

````sh
# bubble sort = O(n^2)
for (i = 0; i < n; i++) {
  for (j = 0; j < n - i - 1 ; j++) {
  }
}

# O(N^2/2) = O(N^2) # dropped constant 2
for (var i = 0; i < n; i++) {
  for (var j = i; j < n; j++) {
  }
}

# O(log N * N) = O(N)
for (int i = N; i > 0; i /= 2) {
  for (int j = 0; j < i; j++) {
  }
}
````

Algorithm X is asymptotically more efficient than Y -
X will always be a better choice for large inputs.

`O` Big O notation - Upper bound.
`Θ` Theta notation - Tight bound.
`Ω` Omega notation - Lower bound.

#### log

````
a^x = b
2^3 = 8

x = log a (b)
3 = log 2 (8)

log 10 (x) = lg (x) = log (x) # by default base = 10 in mathematics
                              # by default base =  2 in programming
ln (x) = log e (x)            # where `e` = 2.72
````
