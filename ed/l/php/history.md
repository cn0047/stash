History of PHP
-

http://php.net/manual/en/appendices.php

#### php 7.2.x

* `object` type hint
* Allow abstract method overriding
* Allow a trailing comma for grouped namespaces `use Foo\Bar\{Foo, Bar, Baz};`

#### php 7.1.x

* **Nullable types.**
* **Void functions.**
* Symmetric array destructuring (like list()).
* Class constant visibility.
* Iterable pseudo-type.
* **Multi catch exception handling.**
* Asynchronous signal handling (no need use tick).
* **HTTP/2 server push support in ext/curl.**

#### php 7.0.x

* **Scalar type declarations.**
* **Return type declarations.**
* Null coalescing operator (??).
* Spaceship operator (<=> returns 0 if both operands are equal, 1 if the left is greater, and -1 if the right is greater).
* Constant arrays using define().
* **Anonymous classes.**
* Closure::call().
* Filtered unserialize() (prevent code injections).
* IntlChar.
* Group use declarations.
* Integer division with intdiv().
* Session options (session_start() now accepts an array).
* random_bytes() and random_int().

#### php 5.6.x

* **Constant scalar expressions.**
* **Variadic functions via ...**
* **Argument unpacking via ...**
* Exponentiation via **
* Use function and use const.
* Phpdbg.
* Default character encoding.
* Files larger than 2 gigabytes in size are now accepted.

#### php 5.5.x

* **Generators.**
* **Try-catch blocks now support a finally.**
* **New password hashing** API.
* Foreach now supports list().
* Empty() supports arbitrary expressions.
* Array and string literal dereferencing.
* Class name resolution via ::class. `ClassName::class`
* **OPcache extension added.**

#### php 5.4.x

* **Traits.**
* **Short array syntax.**
* Function array dereferencing has been added. `foo()[0]`
* Closures now support `$this`.
* Class member access on instantiation has been added. `(new Foo)->bar()`
* Binary number format.

#### php 5.3.x

* **Namespaces.**
* **Late Static Bindings.**
* **Native Closures.**
* Nowdoc syntax is now supported, similar to Heredoc syntax, but with single quotes.
* Constants can now be declared outside a class using the const keyword.
* Class can implement two interfaces that specified a method with the same name.
* $r = ($v) ?: 'No Value';

#### php 5.2.x

#### php 5.1.x
