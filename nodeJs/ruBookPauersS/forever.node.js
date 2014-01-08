/*
npm install forever -g
forever --help
forever start -a -l forever.log -o out.log -e err.log httpserver.js

start                  Запускает сценарий.
stop                   Останавливает выполнение сценария.
stopall                Останавливает выполнение всех сценариев.
restart                Перезапускает сценарий.
restartall             Перезапускает все сценарии, запущенные с помощью Forever.
cleanlogs              Удаляет все записи в журналах.
logs                   Выводит список всех журналов для всех Forever-процессов.
list                   Выводит список всех выполняемых сценариев.
config                 Выводит перечень пользовательских конфигураций.
set <ключ> <значение>  Устанавливает для конфигурации пару ключ-значение.
clear <ключ>           Очищает значение ключа конфигурации.
logs <сценарий|индекс> Закрывает журналы для сценария или индекса.
columns add <столбец>  Добавляет столбец к списку вывода Forever.
columns rm <столбец>   Удаляет столбец для списка вывода Forever.
columns set <столбцы>  Устанавливает все столбцы для списка вывода Forever.
*/

var forever = require('forever');
var child = new (forever.Monitor)('your-filename.js', {
    max: 3,
    silent: true,
    options: []
});
child.on('exit', this.callback);
child.start();