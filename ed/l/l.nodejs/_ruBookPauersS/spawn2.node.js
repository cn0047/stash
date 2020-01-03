// This scritp do the same with: find . -ls | grep test
var spawn = require('child_process').spawn,
find = spawn('find',['.','-ls']),
grep = spawn('grep',['test']);
grep.stdout.setEncoding('utf8');
// направление результатов поиска в адрес grep
find.stdout.on('data', function(data) {
    grep.stdin.write(data);
});
// а теперь запуск grep и вывод результатов
grep.stdout.on('data', function (data) {
    console.log(data);
});
// обработка ошибки для обоих процессов
find.stderr.on('data', function (data) {
    console.log('grep stderr: ' + data);
});
grep.stderr.on('data', function (data) {
    console.log('grep stderr: ' + data);
});
// и завершение обработки для обоих процессов
find.on('exit', function (code) {
    if (code !== 0) {
        console.log('find process exited with code ' + code);
    }
    // продолжение обработки и завершение процесса grep
    grep.stdin.end();
});
grep.on('exit', function (code) {
    if (code !== 0) {
        console.log('grep process exited with code ' + code);
    }
    // продолжение обработки и завершение процесса grep
    grep.stdin.end();
});
