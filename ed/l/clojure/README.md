Clojure
-

[docs](https://clojure.org/guides/getting_started)
[github](https://github.com/clojure/clojure)

Clojure - dynamic and functional dialect of Lisp lang on Java platform.

```clj
inc 5
dec 5

max 1 2 3
min 1 2 3

rem 3 2 ;; remainder of dividing

(println (format "Hello , %s" "World"))

; destructuring
(let [[a b c d] my-vector]
  (println a b c d)
)

```

Predicate - function that evaluates condition and returns true or false.

Macros - used to generate inline code.

Watcher - function added to variable such as atoms and reference variables
which get invoked when value of the variable changes.

#### Data types

* Integers (Decimal , Octal, Hexadecimal, Radix).
* Floating point.
* char.
* Boolean.
* String.
* Nil.
* Atom (shared, synchronous, independent state).

#### Variables types

* short (short number).
* int.
* long.
* float.
* char.
* Boolean.
* String.
