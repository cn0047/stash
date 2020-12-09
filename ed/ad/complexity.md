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
# time complexity:
O(1)      # constant time
O(log(n)) # logarithmic time
O(n)      # linear time
O(n^2)    # n-squared time
O(2^n)    # exponential time
O(n!)     # n-factorial time

# space complexity - same like time complexity ↑

O(n^2 + n^2)            = O(n^2)
O(n^2 + n)              = O(n^2)
O(n + log n)            = O(n) # because log n < n
O(n * log n)            = O(n) # ↑
O(5 * 2^n + 10 * n^100) = O(2^n)
O(n^3 + n^2 + n)        = O(n^3)
O(n^2 + b)              = O(n^2 + b) # because we know nothing about b
O(log 2 n)              = O(log n)
O(n/2)                  = O(n)
O(n^2/2)                = O(n^2)

O(2^n)                  = for fibonacci and for recursion (with creating forks with n levels deep)
O(n!)                   = for recursion / permutations / salesman problems
O(log n)                = if divide in 1/2 or myltiply by 2
                          because: n/2/2/2/2... is n/2^k = 1 is n = 2^k where k = log 2 n
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

# O(n^2/2) = O(n^2) # dropped constant 2
for (var i = 0; i < n; i++) {
  for (var j = i; j < n; j++) {
  }
}

# O(log n * n) = O(n)
for (int i = n; i > 0; i /= 2) {
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
a^x = b           # 2^3 = 8
x = log a (b)     # 3 = log 2 (8)

log 10 (x) = lg (x) = log (x) # by default base = 10 in mathematics
                              # by default base =  2 in programming
ln (x) = log e (x)            # where `e` = 2.72
````
