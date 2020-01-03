let b;

const base64 = new Buffer('some binary data', 'binary').toString('base64');
b = new Buffer(base64, 'base64');
console.log(base64, ' ➡️ ', b.toString());

const ascii = new Buffer('It is test 2', 'base64').toString('ascii');
b = new Buffer(ascii, 'ascii');
console.log(ascii, ' ➡️ ');
