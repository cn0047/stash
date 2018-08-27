Functions
-

````js
a = 1;               typeof a; // "number"
a = 1.1;             typeof a; // "number"
a = '1';             typeof a; // "string"
a = true;            typeof a; // "boolean"
a = {};              typeof a; // "object"
a = [];              typeof a; // "object" But: a instanceof Array; // true
a = new Array(1, 2); typeof a; // "object"
a = null;            typeof a; // "object" And: a instanceof null; //  Uncaught TypeError: Expecting a function in instanceof check, but got null
a = undefined;       typeof a; // "undefined"
a = NaN;             typeof a; // "number"
a = function () {};  typeof a; // "function"
a = /.*/;            typeof a; // "object" But: a instanceof RegExp; // true
a = new RegExp();    typeof a; // "object"
a = new Date();      typeof a; // "object"

function isObject(o) {
    return o instanceof Object && o.constructor === Object;
}

6 + '1'     // 61    - as string.
6 - '1'     // 5     - as int.
"3" + 4 + 5 // "345" - as string.
3 + 4 + "5" // "75"  - as int and as string.

0 == false // true
0 == null // false
false == null // false
null == null // true
NaN == NaN // false
NaN === NaN // false
[] == [] // false
[] === [] // false
{} == {} // false
{} === {} // false

if (true) {} elseif (true) {} // Uncaught SyntaxError: Unexpected token {
if (true) {} else if (true) {} // OK.
````

````js
const t = setTimeout(() => {}, milliseconds);
const i = setInterval(() => {}, milliseconds);
setImmediate(() => {})
// `setImmediate` will always be executed before `setTimeout` and `setInterval` if scheduled within an I/O cycle.
clearTimeout();
clearInterval(i);
````

````js
navigator.cookieEnabled; // Show is cookie allowed.

location.reload();  // Reload document
location.replace(); // Reload document, and don't save action in history

window.history.back
window.history.forward
````

````js
'use strict';

document.getElementById('wrapper').getElementsByClassName('block').getElementsByTagName('img').length
document.querySelectorAll('#wrapper img').length

JSON.stringify(object);                // object to JSON string
JSON.parse(string)                     // JSON string to object

console.group('Application Log', 'font-size: x-large');
console.time('ttt');
console.timeEnd(); // shows spent time
console.log('Application Log');
console.eror('E');
console.debug('D');
console.assert(true);
console.dir(document);
console.groupEnd();

console.log(window.performance.now());
window.performance.mark('start');
window.performance.mark('end');
window.performance.measure('took', 'start', 'end');
console.log(window.performance.getEntriesByType('mark'));
console.log(window.performance.getEntriesByType('measure'));
//available: navigationStart, unloadEventStart, unloadEventEnd, redirectStart, redirectEnd, fetchStart, domainLookupStart, domainLookupEnd, connectStart, connectEnd, secureConnectionStart, requestStart, responseEnd, domLoading, domInteractive, domContentLoadedEventStart, domContentLoadedEventEnd, domComplete, loadEventStart, loadEventEnd

console.log('Code:%s', 200); // Code:200
// %s - string
// %d, %i - integer
// %f - float
// %o - DOM el
// %O - object
// %c - css
````

````js
target.addEventListener(type, listener, options);
document.getElementById('close').addEventListener('click', (e) => {
  e.target.parentElement.style.display = 'none';
});
````

#### Flashback

````js
alert(message);
prompt(message);
confirm(message);

var o = {x: 1, y: 2};
delete o.x; // delete property x
typeof o.x; // undefined
delete o; // false. Cant't delete global var.

o['y']; // Access to associative array.

void varName; // Set undefined to var;

null == undefined  // true
null === undefined // false

typeof val;
tepeof(val);
val instanceof Object;

var d = new Date();
d.constructor == Date; // true

try {
} catch (error) {
} finally {
}

switch (x) {
    case 1:
        break;
    default:
        break;
}

// Convenience method but slow (many optimizations disabled here)...
// It won't create new proberty in scope of object.
// It creates new own execution scope...
with (frames[1].document.forms[0]) {
    name.value = '007'; // Equal to frames[1].document.forms[0].name.value
}
// Better use:
var form = frames[1].document.forms[0];
form.name.value = '007';

el.childNodes;      // child
el.parentNode;      // parent
el.nextSibling;     // next brother
el.previousSibling; // prev brother

function rectangle(w, h) {
    this.with = w;
    this.height = h;
}
rectangle.protytype.area = function () {
    console.log(this.superclass);
}

window.document == self == parent == top
window.defaultStatus = 'status bar string';
window.onerror = function (msg, url, line) {}
````

#### Fetch

````js
// GET
fetch('http://jsonplaceholder.typicode.com/posts')
    .then(res => res.json())
    .then(items => console.log(items))
;
// POST
var data = new FormData();
data.append('userId', 0);
data.append('message', '');
fetch('https://davidwalsh.name/submit-json', {
    method: 'post',
    headers: {
        "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
    },
    body: JSON.stringify({
        email: document.getElementById('email').value,
        answer: document.getElementById('answer').value
    }),
    body: data
});
// Var in url
fetch(`http://myapi.com/posts/${postId}/comments`, {});
````

#### String

````js
str.length;
str.toUpperCase();
str.toLowerCase();
str.IndexOf(str);
str.lastIndexOf(str);
str.slice(7);
str.split();

'James Bond'.match(/(J).*(B)/); // ["James B", "J", "B"]
'James Bond'.replace('Bond', '007'); // James 007
'James Bond James Bond James Bond James Bond'.replace(/Bond/g, '007'); // miltiple replace
````

#### Number

````js
Number(bool);
number.isInteger()
number.isNan();
number.parseFloat()
number.parseInt()
number.NEGATIVE_INFINITY;
number.POSITIVE_INFINITY;

Number(string); parseInt('3 girls'); // 3
parseFloat('3.14 meters');           // 3.14
parseInt('12.34');                   // 12
parseInt('0xFF');                    // 255
parseInt('11', 2);                   // 3
parseInt('ff', 16);                  // 255
parseInt('zz', 36);                  // 1295
parseInt('077', 8);                  // 63
parseInt('077', 10);                 // 77

// Faster than parseInt(inputValue, 10);
// But be careful bitshift operations always return a 32-bit integer.
const val = inputValue >> 0;

Math.round(number); // number rounded to the nearest integer
Math.ceil(number);
Math.floor(number);
Math.random();

String(number);                                        // 12345.6789
number.toString(2);      // to bin                     // 11000000111001.101011011100110001100011111100010100001
number.toString(8);      // to oct                     // 30071.5334614374241
number.toString(16);     // to hex                     // 3039.adcc63f142
number.toFixed();                                      // 12346
number.toFixed(2);       // two characters after comma // 12345.68
number.toExponential(1);                               // 1.2e+4
number.toExponential(3);                               // 1.235e+4
number.toPrecision(4);                                 // 1.235e+4
number.toPrecision(7);                                 // 12345.68
````

#### Boolean

````js
Boolean(bool);
var b = !! bool; // To boolean

const foo = a || b; // a ? a : b; # better use code below
[target = 'default'] = [valueProvidedInFunction] // it works like isset
const bar = !!c; // c ? true : false;
const baz = !c; // c ? false : true;
````

#### Symbol

#### Undefined

````
console.log(typeof v !== 'undefined'); // isset variable v.
````

#### Null

#### Object

````js
Object.assign(dst, src1, src2)
fr_obj = Object.freeze(obj) // freeze obj, but freeze is not recursive! Nested objects won't be frozen.
Object.isFrozen(obj);
Object.seal(obj); // preventing add new properties and marking all existing properties as non-configurable.
Object.defineProperties(obj, props); // defines|modifies object's properties
Object.create // Child.prototype = Object.create(Parent.prototype);

my_obj.hasOwnProperty('cos'); // true
my_obj.keys()
my_obj.values()
my_obj.toString()

var o = {x: 1};
o.toString();                // [object Object]
o.toLocaleString();          // [object Object]
o.valueOf();                 // Object {x: 1}
o.propertyIsEnumerable('x'); // true
o.isPrototypeOf();
o[-1.23] = true;             // Create new property -1.23s
````

Date:

````js
var d = new Date(); // Mon Apr 07 2014 23:09:01 GMT+0300 (EEST)
d.getFullYear();    // 2014
d.getMonth();       // 3
d.getDay();         // 1
d.getHours();       // 24
d.getMinutes();     // 9
d.getSeconds();     // 1

d = Date.parse(str);
````

Errors:

* Error
* EvalError
* InternalError
* RangeError
* ReferenceError
* SyntaxError
* TypeError
* URIError

````
const err = new Error('Name required');
err.status = 400;
err.expose = true;
throw err;
````

Regex:

````
 /hello/.test('hello world');
````

#### Function

````js
f.call(object, 1, 2);
f.apply(object, [1, 2]);
functionWithBoundedContext = f.bind(object, 1, 2);

function f(a, b, c) {
    console.log(b.callee.length);
    console.log(arguments[0]);
    return arguments.length;
}

function f(x) {
    return x * arguments.callee(x-1);
}
````

#### Array

````js
Array.isArray({})
Array.IndexOf(my_array, el);
my_array.IndexOf(el);
my_array.find(callback, this_arg)
my_array.forEach(function(item, index, array) {});
new_array = my_array.filter(callback, this_arg)
new_array = my_array.map(Math.sqrt);
removedItem = my_array.splice(pos, 1); // remove item

var array = new Array(10); // length = 10
array.length;
array.length = 5; // truncate or increase array
array.push('one');
array.pop();
array.unshift('one'); // adds to the beginning of an array
array.shift();
array.reverse();
array.sort(function (a, b) { return a - b; });
array.split();
array.join('.');
array.concat(var_name);

var a = ["dog", "cat", "hen"];
a[100] = "fox";
a.length; // 101
// Remember — the length of the array is one more than the highest index.

for (key in array) {
    console.log(array[key]);
}
var a = [];
a[5] = 5; // Perfectly legal JavaScript that resizes the array.
for (var i=0; i<a.length; i++) {
    // Iterates over numeric indexes from 0 to 5, as everyone expects.
}

var a = [];
a[5] = 5;
for (var x in a) {
    // Shows only the explicitly set index of "5", and ignores 0-4
}

// Somewhere deep in your JavaScript library...
Array.prototype.foo = 1;
// Now you have no idea what the below code will do.
var a = [1,2,3,4,5];
for (var x in a) {
    // Now foo is a part of EVERY array and
    // will show up here as a value of 'x'.
}
````

