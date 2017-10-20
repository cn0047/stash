Complexity of algorithms
-

Drop constants in big O. Not O(2n), not O(3n)...
Just looking for how thing scale goughtly (lineral, quadratic, etc).

For array inside another array - O(a * b),
whre a - 1st array length, b - 2nd.

`time` - it isn't possible to estimate time because it's vary on different machines (CPU, memory, etc.)
`memory`
`count of steps (actions)`

Big O - Dependency between count of iterations and input parameters.

````
O(N^2 + N^2) = O(N^2)
O(N^2 + N) = O(N^2)
O(N + log N) = O(N) (because log N < N)
O(5 * 2^N + 10 * N^100) = O(2^N)
O(N^2 + B) = O(N^2 + B) (because we know nothing about B)
O(log2 N) = O(log N)
O(N/2) = O(N)
O(N^2/2) = O(N^2)
````

If algorithm takes half of elements on iteration than Big O = O(log N).

`for (var i = 0; i < n; i++) for (var j = i; j < n; j++) {}` - O(N^2/2) = O(N^2)
