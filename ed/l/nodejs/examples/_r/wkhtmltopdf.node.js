var spawn = require('./child_process').spawn;
// аргументы командной строки
var url = process.argv[2];
var output = process.argv[3];
if (url && output) {
    var wkhtmltopdf = spawn('wkhtmltopdf.sh', [url, output]);
    wkhtmltopdf.stdout.setEncoding('utf8');
    wkhtmltopdf.stdout.on('data', function (data) {
        console.log(data);
    });
    wkhtmltopdf.stderr.on('data', function (data) {
        console.log('stderr: ' + data);
    });
    wkhtmltopdf.on('exit', function (code) {
        console.log('child process exited with code ' + code);
    });
} else {
    console.log('You need to provide a URL and output file name');
}



var pdftk = spawn('pdftk', [__dirname + '/pdfs/datasheet-node.pdf', 'dump_data']);
pdftk.stdout.on('data', function (data) {
    // преобразование результатов в объект
    var array = data.toString().split('\n');
    var obj = {};
    array.forEach(function(line) {
        var tmp = line.split(':');
        obj[tmp[0]] = tmp[1];
    });
    // вывод количества страниц
    console.log(obj['NumberOfPages']);
});
pdftk.stderr.on('data', function (data) {
    console.log('stderr: ' + data);
});
pdftk.on('exit', function (code) {
    console.log('child process exited with code ' + code);
});