for (var i = 1; i <= 5; i++) {
    setTimeout(function() {
        console.log(i);
    }, 1000);
}

/*
Result:
6
6
6
6
6
*/

for (var i = 1; i <= 5; i++) {
    (function(i) {
        setTimeout(function() {
            console.log(i);
        }, 1000);
    })(i);
}
// OR
for (let i = 1; i <= 5; i++) {
    setTimeout(function() {
        console.log(i);
    }, 1000);
}

/*
Result:
1
2
3
4
5
*/
