Code style
-

[source](https://github.com/felixge/node-style-guide)
[airbnb](https://github.com/airbnb/javascript)

## Formatting

* 2 Spaces for indentation

* Use UNIX-style newlines

* No trailing whitespace

* Use Semicolons

* 80 characters per line

* Use single quotes

* Opening braces go on the same line

````
if (true) {
  console.log('winning');
}
````

* Declare one variable per var statement

````
var keys   = ['foo', 'bar'];
var values = [23, 42];
````

## Naming Conventions

* Use lowerCamelCase for variables, properties and function names

* Use UpperCamelCase for class names

* Use UPPERCASE for Constants

## Variables

* Object / Array creation

````
var a = [
    'hello',
    'world',
];
var b = {
  good: 'code',
  'is generally': 'pretty',
};
````

## Conditionals

* Use the === operator

## Functions

* Write small functions

Limit yourself to ~15 lines of code per function...

* Return early from functions

````
function isPercentage(val) {
  if (val < 0) {
    return false;
  }
  if (val > 100) {
    return false;
  }
  return true;
}
````

* Name your closures

Will produce better stack traces, heap and cpu profiles...
You can reuse this name for recursion.
It is also description for code.

````
req.on('end', function onEnd() {
  console.log('winning');
});
````

* No nested closures

* Method chaining

One method per line should be used if you want to chain methods.

````
User
  .findOne({ name: 'foo' })
  .populate('bar')
  .exec(function(err, user) {
    return true;
  });
````

## Comments

* Use slashes for comments

## Miscellaneous

* Requires at top

* Do not extend built-in prototypes
