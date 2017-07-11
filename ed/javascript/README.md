js
-

A polyfill is a browser fallback, made in JavaScript,
that allows functionality you expect to work in modern browsers to work in older browsers, e.g.,
to support canvas (an HTML5 feature) in older browsers.

A shim is more generalized. A polyfill is a type of shim.

Variables can be hoisted. Hoisted means, declare them all on top of the function.

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

// Convenience method but slow
with (frames[1].document.forms[0]) {
    name.value = '007'; // Equal to frames[1].document.forms[0].name.value
}
// Better use:
var form = frames[1].document.forms[0];
form.name.value = '007';

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
    * Undefined
* Object
    * Null (a = null; typeof a; // "object")
    * Object
    * Function (it isn't object because pass by value not by link, @see examples/bubbleSort.js)

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
````
