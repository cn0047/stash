var util = require('util');
// var jsdom = require('jsdom');
// console.log(util.inspect(jsdom, true, null, true));

// определение исходного объекта
function first() {
    // В JavaScript this представляет собой контекст объекта, который может меняться.
    // Единственным способом сохранить данные для метода объекта является присва-
    // ивание this переменной объекта, в данном случае — self, а затем использовать
    // переменную внутри любых функций объекта.
    // Запуск этого приложения приведет к следующ
    var self = this;
    this.name = 'first';
    this.test = function() {
        console.log(self.name);
    };
}
first.prototype.output = function() {
    console.log(this.name);
}
// наследование из first
function second() {
    // Если убрать эту строку из конструктора второго объекта, любой вызов метода
    // output в отношении второго объекта будет успешным, а вот вызов метода test
    // приведет к ошибке и заставит Node-приложение завершить работу с сообщением
    // о том, что переменная test не определена.
    // Метод call устанавливает связь конструкторов двух объектов, гарантируя вызов
    // суперконструктора наряду с конструктором. Суперконструктор является конс-
    // труктором для наследуемого объекта.
    second.super_.call(this);
    /*
    exports.inherits = function(ctor, superCtor) {
        ctor.super_ = superCtor;
        ctor.prototype = Object.create(superCtor.prototype, {
            constructor: {
                value: ctor,
                enumerable: false,
                writable: true,
                configurable: true
            }
        });
    };
    */
    this.name = 'second';
}
// При вызове util.inherits фрагмент super_ в качестве свойства присваивается объекту second
util.inherits(second,first);
var two = new second();
function third(func) {
    this.name = 'third';
    this.callMethod = func;
}
var three = new third(two.test);
// при всех трех вызовах должно быть выведено "second"
two.output();
two.test();
three.callMethod();
