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
