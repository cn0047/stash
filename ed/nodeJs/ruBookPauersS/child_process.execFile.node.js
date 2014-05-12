var execfile = require('child_process').execFile, child;
child = execfile('./app.js', function(error, stdout, stderr) {
    if (error == null) {
        console.log('stdout: ' + stdout);
    }
});
