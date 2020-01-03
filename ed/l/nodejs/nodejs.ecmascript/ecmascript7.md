ECMAScript 7 (2016)
-

````js
// Object Observe
var obj = {};
Object.observe(obj,function(changes) {console.log(changes);});
obj.name = "hemanth"; // [ { type: 'new', object: { name: 'hemanth' }, name: 'name' } ]

// Object.getOwnPropertyDescriptors
let life = { answer: 42 };
Object.getOwnPropertyDescriptor(life, 'answer');
// { configurable: true, enumerable: true, value: 42, writable: true }

// Array.prototype.includes
[1, 2, NaN].includes(NaN); // true
[0,+1,-1].includes(42); // false

// Typed Objects

// Exponentiation Operator
let cubed = x => x ** 3;
cubed(2) // 8;

// Array comprehensions
[for (num of numbers) Math.sqrt(num)]; // => [ 1, 2, 3 ]

// Generator comprehensions
(for (i of [ 2, 4, 6 ]) i*i ); // generator function which yields 4, 16, and 36

// Async functions

// Async generators
````
