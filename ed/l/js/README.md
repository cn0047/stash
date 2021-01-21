JavaScript
-
Since 1995.

JS - is a high-level (not assembler), dynamic (behavior determines in runtime)
weakly typed, object-based, multi-paradigm (event-driven, functional, imperative, prototype-based, object-oriented)
and interpreted client-side programming language.

A polyfill is a browser fallback, made in JavaScript,
that allows functionality you expect to work in modern browsers to work in older browsers, e.g.,
to support canvas (an HTML5 feature) in older browsers.

A shim is more generalized. A polyfill is a type of shim.

Lexical scope - code author's scope (context of the surrounding code)...
To cheat with lexical scope use `eval` or `with`.
`eval` is slow because no optimizaton for unknown code.

Dynamical scope - Runtime scope.

Create new scope - use IIFE or catch block.

IIFE (Immediately-Invoked Function Expression) - needed for creating new scope,
and hide some stuff from outer scope.

**Hoisting**. Variable can be used before it has been declared.
Hoisting is JS's default behavior of moving all declarations to the top of the current scope.
JS only hoists declarations, not initializations (definitions or expressions).
Functions hoisted before variables.
Hoisting provides ability to create mutual-recursion.

**Temporal dead zone** - use block-scoped variable (let keyword) before `let` occurs.
`let`, `class` don't hoists.

**Lambda** - function that is used as data (value) (assigned to variable or passed between functions).
Parameter for another function, return value of a function etc.

**Closure** - if you define a function inside another function,
the inner function will have full access to all the variables
that are declared and available in the outer function's scope.
If the outer functions completes execution and returns -
the inner function will still have access to all of the variables
that were part of the outer function when the inner function was returned.

Callback hell - IoC. We lose control of our code and rely on code which will call our callback,
but how it will be called, how much time, etc - it is out of our control.

**this**, The value of this within any given function call is determined by *how the function is called*.
Function's own execution context.

Rule 4: `new` keyword.

Rule 3: Explicit binding - `func.call` & `func.apply`;

Rule 2: Implicit binding - reference to a function via object property reference (`obj.foo()`);

Rule 1: Default binding - strict mode - undefined; non strict - global object (window);

**new**.

1. Creates new object.

2. Links new object to constructor function (`prototype`).

3. Makes `this` variable point to the new object.

4. Executes constructor function using the new object and implicit perform `return this`;

5. Assigns constructor function name to new object's property `constructor`.

`Object.create` - performs first 2 steps form `new`.

`new` override hard binding.

**Generators** - function executions that can be suspended and resumed at a later point.
`function*` & `yield`.
Generator function on 1st call returns iterator.
`yield` in function will pause iterator, and `generatorFunctionIterator.next()` wil resume.
`yield` - message pass mechanizm (in generator and out from generator).

**Prototype** helps objects to be linked together in a hierarchy.
Each function in JavaScript has a member called prototype,
which is responsible for providing values when an object is asked for them.
Adding a property to the prototype of a function object will make it available at the constructed object.
When an object is asked for a property that it does not have,
its parent object will be asked.

__proto__ - internal property used by the js engine for inheritance.

Shadowing - it is when parent class* have method which is overridden in child class*,
so child can't call parent class and have to use: `Parent.prototype.func.call(this)`.

````js
function f () {}
f.prototype.p = 200;

console.log(f.p); // undefined

f.prototype.getP = function () {
    return this.p;
}

console.log(f.getP()); // Uncaught TypeError: f.getP is not a function(…)

var i = new f();
console.log(i.p); // 200
console.log(i.getP()); // 200
````
````js
function f () {}
f.prototype.p = 200;
f.prototype.ff = function () {};
f.prototype.ff.prototype.pp = 204;

i = new f();
console.log(i.p); // 200
console.log(i.ff.pp); // undefined

i.ff = new f();
console.log(i.ff.pp); // undefined
console.log(i.ff.p); // 200
````

#### Events

Phases (event flow) (`var phase = event.eventPhase;`):

* CAPTURING_PHASE
  Event propagated through the target's ancestor objects: Window -> Document -> HTMLHtmlElement

* AT_TARGET
  The event has arrived at the event's target.
  Event listeners registered for this phase are called at this time. 
  If `Event.bubbles` (`event.stopPropagation()`) is false, processing the event is finished.

* BUBBLING_PHASE
  The event is propagating back up through the target's ancestors in reverse order,
  starting with the parent, and eventually reaching the containing Window.

Page loading event:

* `window.onload`

Fires at the end of the document loading process.
At this point, all of the objects in the document are in the DOM,
and all the images, scripts, links and sub-frames have finished loading.

* `document.onload`

Called when the DOM tree is completed.

#### Data Types

* Primitives:
    * String.
    * Number.
    * Boolean.
    * Symbol (new in ECMAScript 6).
* Special:
    * Null (`a = null; typeof a; // "object"`).
    * Undefined.
* Object:
    * Object.
    * Array.
    * Function (it isn't object because pass by value not by link, @see examples/function.js).
    * etc.
