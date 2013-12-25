var vm = require('vm');
var util = require('util');
var obj = { name: 'Shelley', domain: 'burningbird.net'};
// компиляция сценария
var script_obj = vm.createScript("var str = 'My name is ' + name + ' at ' + domain", 'test.vm');
// запуск в новом контексте
script_obj.runInNewContext(obj);
// обследование объекта песочницы
console.log(util.inspect(obj));

var fs = require('fs');
fs.readFile('suspicious.js', 'utf8', function(err, data) {
    if (err) return console.log(err);
    try {
        console.log(data);
        var obj = { name: 'Shelley', domain: 'burningbird.net'};
        // компилирование сценария
        var script_obj = vm.createScript(data, 'test.vm');
        // запуск в новом контексте
        script_obj.runInNewContext(obj);
        // проверка объекта песочницы
        console.log(util.inspect(obj));
    } catch(e) {
        console.log(e);
    }
});

fs.readFile('suspicious.js', 'utf8', function(err, data) {
    if (err) return console.log(err);
    try {
        var obj = { name: 'Shelley', domain: 'burningbird.net' };
        // компиляция сценария
        var script_obj = vm.createScript(data, 'test.vm');
        // создание контекста
        var ctx = vm.createContext(obj);
        // запуск в новом контексте
        script_obj.runInContext(ctx);
        // обследование объекта
        console.log(util.inspect(obj));
        // обследование контекста
        console.log(util.inspect(ctx));
    } catch(e) {
        console.log(e);
    }
});