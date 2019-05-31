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
O(N^2 + N^2)            = O(N^2)
O(N^2 + N)              = O(N^2)
O(N + log N)            = O(N) # because log N < N
O(5 * 2^N + 10 * N^100) = O(2^N)
O(N^3 + N^2 + N)        = O(N^3)
O(N^2 + B)              = O(N^2 + B) # because we know nothing about B
O(log2 N)               = O(log N)
O(N/2)                  = O(N)
O(N^2/2)                = O(N^2)

O(2^N)                  = for fibonacci
O(N!)                   = for recursion
O(N!)                   = salesman problem
O(log N)                = if divide in 1/2 or myltiply by 2

N # greater than log N
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

# O(N^2/2) = O(N^2)
for (var i = 0; i < n; i++) {
  for (var j = i; j < n; j++) {
  }
}

# O(N)
for (int i = N; i > 0; i /= 2) {
  for (int j = 0; j < i; j++) {
  }
}
````

Algorithm X is asymptotically more efficient than Y -
X will always be a better choice for large inputs.

Big O notation - Upper bound.
Omega notation - Lower bound.
Theta notation - Tight bound.
