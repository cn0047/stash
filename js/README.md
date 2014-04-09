js
-

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
array.push('one');
array.unshift('one');
array.pop();
array.shift();
array.splice(3, 2);
array.reverse();
array.sort();
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

String(number);                                       // 12345.6789
number.toString(2);      // to bin                    // 11000000111001.101011011100110001100011111100010100001
number.toString(8);      // to oct                    // 30071.5334614374241
number.toString(16);     // to hex                    // 3039.adcc63f142
number.toFixed();                                     // 12346
number.toFixed(2);       // to characters after comma // 12345.68
number.toExponential(1);                              // 1.2e+4
number.toExponential(3);                              // 1.235e+4
number.toPrecision(4);                                // 1.235e+4
number.toPrecision(7);                                // 12345.68
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
void varName; // Set undefined to var;

null == undefined  // true
null === undefined // false

typeof val;
tepeof(val);
val instanceof Object;

````