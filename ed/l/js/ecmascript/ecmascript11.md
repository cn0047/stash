ECMAScript 11 (2020)
-

````js
// Optional chaining
console.log(smartphones.companies?.motorola)
// output is: undefined
// NOT: TypeError: Cannot read property 'motorola' of undefined

console.log(smartphones?.[1])
// output is: undefined
// NOT: TypeError: Cannot read property '1' of null

console.log(phoneApple?.())
// undefined
// NOT: TypeError: phoneApple is not a function

// Nullish Coalescing
let number = theNumber ?? 5

// Private Fields
class Smartphones {
  #phone_color = "silver";
}

// Static Fields
class Smartphone {
  static create_smartphone(color) {
  }
}

// Top Level Await
const color = await fetch(silver) // no need to wrap into async func

// Promise.allSettled
Promise.allSettled([promise_1, promise_2]).then()

// Dynamic Import

// Regex MatchAll

// globalThis
window == globalThis // true

// BigInt
````
