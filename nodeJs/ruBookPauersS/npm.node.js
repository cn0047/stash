npm help npm

npm install jquery

# Если нужна глобальная установка пакета, следует воспользоваться ключом -g или ‐‐global:
npm -g install connect

npm uninstall modulename

# обновить отдельно взятый модуль:
npm update modulename

# список установленных пакетов:
npm ls

# модули установлены глобально:
npm ls -g

# список конфигурационных параметров npm:
npm config list
npm config ls -l
npm config delete keyname
npm config set keyname value

/*
Modules:
Underscore:      Предоставляет полезные JavaScript-функции общего назначения.
Coffee-script:   Позволяет использовать язык CoffeeScript, компилируемый в JavaScript.
Request:         Упрощенный клиент HTTP-запросов.
Express:         Инфраструктура.
Optimist:        Предлагает упрощенный синтаксический разбор ключей.
Async:           Предоставляет функции и схемы для асинхронного кода.
Connect:         Связующее программное обеспечение.
Colors:          Добавляет цвета на консоли.
Uglify-js:       Парсер и компрессор-форматировщик.
Socket.IO:       Позволяет вести обмен данными между клиентом и сервером в реальном времени.
Redis:           Клиент Redis.
Jade:            Движок шаблонов.
Commander:       Модуль для программ командной строки.
Mime:            Предлагает поддержку MIME-расширений файлов.
JSDOM:           Реализует W3C DOM.
*/


# Colors
npm install colors

var colors = require('colors');
console.log('This Node kicks it!'.rainbow.underline);
console.log('We be Nodin'.zebra.bold);
console.log('rainbow'.rainbow, 'zebra'.zebra);
colors.setTheme({
    mod1_warn: 'cyan',
    mod1_error: 'red',
    mod2_note: 'yellow'
});
console.log("This is a helpful message".mod2_note);
console.log("This is a bad message".mod1_error);


# Optimist
var argv = require('optimist').argv;
console.log(argv.o + " " + argv.t);

# Следующая команда приведет к выводу на консоль значений 1 и 2
./app.js -o 1 -t 2

console.log(argv.one + " " + argv.two);
./app2.js --one="My" --two="Name"
