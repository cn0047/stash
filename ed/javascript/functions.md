Functions
-

#### String

````
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

````
Number(bool);
number.isInteger()
number.isNan();
number.NEGATIVE_INFINITY;
number.parseFloat()
number.parseInt()
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

Math.round(number);
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

````
Boolean(bool);
var b = !! bool; // To boolean
````

#### Symbol

#### Undefined

#### Null

#### Object

````
Object.assign(dst, src1, src2)
my_obj.hasOwnProperty('cos'); // true
my_obj.keys()
my_obj.values()
my_obj.toString()
````

#### Function

````
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
````

#### Array

````
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
array.unshift('one'); // adds to the beginning of an array
array.pop();
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
````

