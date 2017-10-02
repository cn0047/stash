Lodash
-
4.17.4

#### Collection

````js
# flatMap
function duplicate(n) {
  return [n, n];
}
_.flatMap([1, 2], duplicate);
// => [1, 1, 2, 2]

# invokeMap
_.invokeMap([[5, 1, 7], [3, 2, 1]], 'sort');
// => [[1, 5, 7], [1, 2, 3]]
````

reject - like `filter` but condition must return false.
sample - random element from collection.

#### Function

ary - func for n arguments and ignore else arguments.
debounce - delays func invoking.
flip - func invokes func with reversed arguments.
memoize - fun that memoizes the result of func.
spread - func apply.
throttle - invokes func at most once per every wait milliseconds.
unary - func accepts up to one argument and ignoring else.

#### Lang

conformsTo - checks if object conforms to source.
isNative - is native js func.

#### Number

clamp - pick closesv value from range.

#### Object

invert - flip keys and values.
merge - like `assign`.
result - get value by path.
toPairs - creates array of own enumerable string keyed-value pairs.
transform - alternative to `reduce`.
update - lik `set`.

#### Seq

tap - inject funct in chain of functions.
thru - like `tap`.

#### String

template - compiled template func.

#### Util

attempt - attempts to invoke func and result or the caught error object.
cond - func with pairs and invoke func for matched key.
conforms - func which check if objects match.
identity - returns the first argument it receives.
matches - partial deep comparison.
method - invokes method at path.
noop - returns undefined.
property - value by path.
times - invokes the iteratee n times.
