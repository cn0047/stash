httpProxy.createServer(function(req,res,proxy) {
debugger;
if (req.url.match(/^\/node\//))
    proxy.proxyRequest(req, res, {host: 'localhost', port: 8000 });
else
    proxy.proxyRequest(req,res, {host: 'localhost', port: 8124 });
}).listen(9000);

/*
node debug debugger.node.js
Для пе рехода к следующей контрольной точке нужно ввести команду cont или ее аббреви атуру c .
Вы можете выполнять код в пошаговом режиме, используя команду next ( n )
Для перехода к следующей инструкции, команду step ( s )
Для перехода к следующей инструкции с заходом в код функций или команду out ( o )
Вы также можете установить новую контрольную точку либо на текущей строке с помощью команды setBreakpoint ( sb )
Снимается контрольная точка командой clearBreakpoint ( cb ).
Команда backtrace пригодится для обратной трассировки (вывода списка текущих активных вызовов функций) выполняемого фрагмента кода
Когда вам понадобится просмотреть перечень доступных команд, наберите команду help

npm install -g node-inspector
node --debug app.js
node-inspector
Go to browser and debug at browser fireBug.

var assert = require('assert');

npm install zombie

npm install nodeload -g

npm install nodemon
*/