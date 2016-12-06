// npm install step
var fs = require('fs'),
Step = require('step');
try {
    Step (
        function readData() {
            fs.readFile('./data/data1.txt', 'utf8', this);
        },
        function modify(err, text) {
            if (err) throw err;
            return text.replace(/somecompany\.com/g,'burningbird.net');
        },
        function writeData(err, text) {
            if (err) throw err;
            fs.writeFile('./data/data1.txt', text, this);
        }
    );
} catch(err) {
    console.error(err);
}