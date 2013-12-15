var fs = require('fs');
var writeStream = fs.createWriteStream('./log.txt', {
    'flags' : 'a',
    'encoding' : 'utf8',
    'mode' : 0666
});
try {
    // получение списка файлов
    fs.readdir('./data/', function(err, files) {
        // для каждого файла
        files.forEach(function(name) {
            // изменение контента
            fs.readFile('./data/' + name,'utf8', function(err,data) {
                if (err) throw err;
                var adjData = data.replace(/somecompany\.com/g,'burningbird.net');
                // запись в файл
                fs.writeFile('./data/' + name, adjData, function(err) {
                    if (err) throw err;
                    // запись в журнал
                    writeStream.write('changed ' + name + '\n', 'utf8', function(err) {
                        if(err) throw err;
                    });
                });
            });
        });
    });
} catch(err) {
console.error(util.inspect(err));
}