var fs = require('fs'),
    Step = require('step'),
    files,
    _dir = './data/';
try {
    Step (
        function readDir() {
            fs.readdir(_dir, this);
        },
        function readFile(err, results) {
            if (err) throw err;
            files = results;
            var group = this.group();
            results.forEach(function(name) {
                fs.readFile(_dir + name, 'utf8', group());
            });
        },
        function writeAll(err, data) {
            if (err) throw err;
            for (var i = 0; i < files.length; i++) {
                var adjdata = data[i].replace(/somecompany\.com/g,'burningbird.net');
                fs.writeFile(_dir + files[i], adjdata, 'utf8',this);
            }
        }
    );
} catch(err) {
    console.log(err);
}