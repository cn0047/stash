var fs = require('fs');
fs.open('fp.txt', 'w', 0666, function(error, fp) {
    fs.write(fp, 'Hello world!', null, 'utf-8', function() {
        fs.close(fp, function(error) {});
    });
});