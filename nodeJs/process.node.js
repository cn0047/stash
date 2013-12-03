console.log(process.execPath);
console.log(process.version);
console.log(process.platform);

// Add this code after: >nodejs -i
process.stdin.resume();
process.stdin.on('data', function (chunk) {
    process.stdout.write('data: ' + chunk);
});
// Type anything to terminal, it right away will be printed.

// nextTick() function call accepted callback function not right away,
// but at next executing asynchFunction().
function asynchFunction = function (data, callback) {
    process.nextTick(function() {
        callback(val);
    });
);
