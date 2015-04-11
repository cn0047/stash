js
-

````js
JSON.stringify(object);                // object to JSON
JSON.parse(string)                     // string to JSON

console.log(typeof v !== 'undefined'); // isset variable v.
````

####Flashback
````js
for (key in array) {
    console.log(array[key]);
}

alert(message);
prompt(message);
confirm(message);

Array.IndexOf(array, el);

var array = new Array(10); // length = 10
array.length;
array.length = 5; // truncate or increase array
array.push('one');
array.unshift('one');
array.pop();
array.shift();
array.splice(3, 2);
array.reverse();
array.sort();
array.sort(function (a, b) { return a - b; });
array.split();
array.join('.');
array.concat(varName);

str.length;
str.toUpperCase();
str.toLowerCase();
str.IndexOf(str);
str.lastIndexOf(str);
str.slice(7);

'James Bond'.match(/(J).*(B)/); // ["James B", "J", "B"]
'James Bond'.replace('Bond', '007'); // James 007

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

// Convenience method bu slow
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

####Data Types
* Primitives:
    * String.
    * Number.
    * Boolean.
    * Symbol (new in ECMAScript 6).
* Object:
    * Object.
    * Array.
* Special:
    * Null.
    * Undefined.

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

if (true) {} elseif (true) {} // Uncaught SyntaxError: Unexpected token {
if (true) {} else if (true) {} // OK.

console.log(
    [1, 2, 3].map(function (v) {
        return 'value:'+v;
    })
);
// ["value:1", "value:2", "value:3"]
````
