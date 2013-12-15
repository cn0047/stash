function test_and_load(filename) {
    var promise = new process.Promise();
    fs.stat(filename).addCallback(function (stat) {
        // фильтрация элементов, не являющихся файлами
        if (!stat.isFile()) { promise.emitSuccess(); return; }
        // в противном случае - чтение файла
        fs.readFile(filename).addCallback(function (data) {
            promise.emitSuccess(data);
        }).addErrback(function (error) {
            promise.emitError(error);
        });
    }).addErrback(function (error) {
        promise.emitError(error);
    });
    return promise;
}



var File = require('file');
var promise = File.read('mydata.txt');
promise.addCallback(function (data) {
    // обработка данных
});