function Universe() {
    // сохраненный экземпляр
    var instance = this;
    // создать новый экземпляр
    this.start_time = 0;
    this.bang = “Big”;
    // переопределить конструктор
    Universe = function () {
        return instance;
    };
}
// проверка
var uni = new Universe();
var uni2 = new Universe();
uni === uni2; // true

// добавить свойство в прототип
Universe.prototype.nothing = true;
var uni = new Universe();
// добавить еще одно свойство в прототип
// уже после создания первого объекта
Universe.prototype.everything = true;
var uni2 = new Universe();
// Проверка:
// объект имеет доступ только
// к оригинальному прототипу
uni.nothing;        // true
uni2.nothing;       // true
uni.everything;     // undefined
uni2.everything;    // undefined
// это выражение дает ожидаемый результат:
uni.constructor.name; // “Universe”
// а это нет:
uni.constructor === Universe; // false
