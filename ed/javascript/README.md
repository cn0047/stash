js
-

````js
'use strict';

document.getElementById('wrapper').getElementsByClassName('block').getElementsByTagName('img').length
document.querySelectorAll('#wrapper img').length

console.group('Application Log');
console.log('Application Log');
console.groupEnd();

JSON.stringify(object);                // object to JSON string
JSON.parse(string)                     // JSON string to object

console.log(typeof v !== 'undefined'); // isset variable v.

console.log('Code:%s', 200); // Code:200
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

#### Prototype

Prototype helps objects to be linked together in a hierarchy.
Each function in JavaScript has a member called prototype,
which is responsible for providing values when an object is asked for them.
Adding a property to the prototype of a function object will make it available at the constructed object.
When an object is asked for a property that it does not have,
its parent object will be asked.

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

#### Flashback

````js
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
for (var x in a){
    // Now foo is a part of EVERY array and
    // will show up here as a value of 'x'.
}

alert(message);
prompt(message);
confirm(message);

Array.IndexOf(array, el);

var array = new Array(10); // length = 10
array.length;
array.length = 5; // truncate or increase array
array.push('one');
array.unshift('one'); // adds to the beginning of an array
array.pop();
array.shift();
array.splice(3, 2);
array.reverse();
array.sort();
array.sort(function (a, b) { return a - b; });
array.split();
array.join('.');
array.concat(varName);

var a = ["dog", "cat", "hen"];
a[100] = "fox";
a.length; // 101
// Remember — the length of the array is one more than the highest index.

str.length;
str.toUpperCase();
str.toLowerCase();
str.IndexOf(str);
str.lastIndexOf(str);
str.slice(7);

'James Bond'.match(/(J).*(B)/); // ["James B", "J", "B"]
'James Bond'.replace('Bond', '007'); // James 007
'James Bond James Bond James Bond James Bond'.replace(/Bond/g, '007'); // miltiple replace

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
number.isNan();

Number(bool);
Number(string); parseInt('3 girls'); // 3
parseFloat('3.14 meters');           // 3.14
parseInt('12.34');                   // 12
parseInt('0xFF');                    // 255
parseInt('11', 2);                   // 3
parseInt('ff', 16);                  // 255
parseInt('zz', 36);                  // 1295
parseInt('077', 8);                  // 63
parseInt('077', 10);                 // 77
Number.NEGATIVE_INFINITY;

Math.round(number);
Math.ceil(number);
Math.floor(number);
Math.random();

var d = new Date(); // Mon Apr 07 2014 23:09:01 GMT+0300 (EEST)
d.getFullYear();    // 2014
d.getMonth();       // 3
d.getDay();         // 1
d.getHours();       // 24
d.getMinutes();     // 9
d.getSeconds();     // 1

Boolean(bool);
var b = !! bool; // To boolean

var o = {x: 1, y: 2};
delete o.x; // delete property x
typeof o.x; // undefined
delete o; // false. Cant't delte global var.

o['y']; // Access to associative array.

void varName; // Set undefined to var;

null == undefined  // true
null === undefined // false

typeof val;
tepeof(val);
val instanceof Object;

var d = new Date();
d.constructor == Date; // true

var o = {x: 1};
o.toString();                // [object Object]
o.toLocaleString();          // [object Object]
o.valueOf();                 // Object {x: 1}
o.propertyIsEnumerable('x'); // true
o.isPrototypeOf();
o[-1.23] = true;             // Create new property -1.23

Math.hasOwnProperty('cos'); // true

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

// Convenience method but slow
with (frames[1].document.forms[0]) {
    name.value = '007'; // Equal to frames[1].document.forms[0].name.value
}
// Better use:
var form = frames[1].document.forms[0];
form.name.value = '007';

["dog", "cat", "hen"].forEach(function(currentValue, index, array) {
  console.log(currentValue);
});

setTimeout("alert(200);", 3000);
clearTimeout();
var i = setInterval("alert(200);", 3000);
clearInterval(i);

navigator.cookieEnabled; // Show is cookie allowed.

location.reload();  // Reload document
location.replace(); // Reload document, and don't save action in history

el.childNodes;      // child
el.parentNode;      // parent
el.nextSibling;     // next brother
el.previousSibling; // prev brother

f.bind(object, 1, 2);
f.call(object, 1, 2);
f.apply(object, [1, 2]);

function f(a, b, c) {
    console.log(b.callee.length);
    console.log(arguments[0]);
    return arguments.length;
}

function f(x) {
    return x * arguments.callee(x-1);
}

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

#### Data Types

* Primitives:
    * String
    * Number
    * Boolean
    * Symbol (new in ECMAScript 6)
    * Null
    * Undefined
* Object

OR

* Primitives:
    * String.
    * Number.
    * Boolean.
    * Symbol (new in ECMAScript 6).
* Special:
    * Null.
    * Undefined.
* Object:
    * Object.
    * Array.
    * Function.
    * etc.

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

console.log(
    [1, 2, 3].map(function (v) {
        return 'value:'+v;
    })
);
// ["value:1", "value:2", "value:3"]
````
