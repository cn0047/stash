var fs = require('fs');
var path = require('path');
fs.realpath(__dirname+'/'+'data.txt', function(err, pathname) {
    var p = {
        'dirname': path.dirname(pathname),
        'basename': path.basename(pathname),
        'extension': path.extname(pathname).substring(1),
        'filename': path.basename(pathname, path.extname(pathname))
    };
    console.log('The dirname is '+p['dirname']);
    console.log('The basename is '+p['basename']);
    console.log('The extension is '+p['extension']);
    console.log('The filename is '+p['filename']);
});

fs.chmod(__dirname+'/'+'data.txt', 0640, function() {
    fs.chmod(__dirname+'/'+'data.txt', '640', function() {
    });
});