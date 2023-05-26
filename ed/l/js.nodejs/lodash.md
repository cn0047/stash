Lodash
-
4.17.4

#### Collection

````js
// flatMap
function duplicate(n) {
  return [n, n];
}
_.flatMap([1, 2], duplicate);
// => [1, 1, 2, 2]

// invokeMap
_.invokeMap([[5, 1, 7], [3, 2, 1]], 'sort');
// => [[1, 5, 7], [1, 2, 3]]
````

<br>`reject` - like `filter` but condition must return false.
<br>`sample` - random element from collection.

#### Function

<br>`ary` - func for n arguments and ignore else arguments.
<br>`debounce` - delays func invoking.
<br>`flip` - func invokes func with reversed arguments.
<br>`memoize` - fun that memoizes the result of func.
<br>`spread` - func apply.
<br>`throttle` - invokes func at most once per every wait milliseconds.
<br>`unary` - func accepts up to one argument and ignoring else.

#### Lang

<br>`conformsTo` - checks if object conforms to source.
<br>`isNative` - is native js func.

#### Number

<br>`clamp` - pick closesv value from range.

#### Object

<br>`invert` - flip keys and values.
<br>`merge` - like `assign`.
<br>`result` - get value by path.
<br>`toPairs` - creates array of own enumerable string keyed-value pairs.
<br>`transform` - alternative to `reduce`.
<br>`update` - lik `set`.

#### Seq

<br>`tap` - inject funct in chain of functions.
<br>`thru` - like `tap`.

#### String

<br>`template` - compiled template func.

#### Util

<br>`attempt` - attempts to invoke func and result or the caught error object.
<br>`cond` - func with pairs and invoke func for matched key.
<br>`conforms` - func which check if objects match.
<br>`identity` - returns the first argument it receives.
<br>`matches` - partial deep comparison.
<br>`method` - invokes method at path.
<br>`noop` - returns undefined.
<br>`property` - value by path.
<br>`times` - invokes the iteratee n times.
