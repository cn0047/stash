function Universe() {
    // сохраненный экземпляр
    var instance;
    // переопределить конструктор
    Universe = function Universe() {
        return instance;
    };
    // перенести свойства прототипа
    Universe.prototype = this;
    // создать экземпляр
    instance = new Universe();
    // переустановить указатель на конструктор
    instance.constructor = Universe;
    // добавить остальную функциональность
    instance.start_time = 0;
    instance.bang = “Big”;
    return instance;
}

// добавить свойство в прототип и создать экземпляр
Universe.prototype.nothing = true; // true
var uni = new Universe();
Universe.prototype.everything = true; // true
var uni2 = new Universe();
// тот же самый экземпляр
uni === uni2; // true
// все свойства прототипа доступны
// независимо от того, когда они были добавлены
uni.nothing && uni.everything && uni2.nothing && uni2.everything; // true
// обычные свойства объекта также доступны
uni.bang; // “Big”
// ссылка на конструктор содержит правильный указатель
uni.constructor === Universe; // true


var Universe;
(function () {
    var instance;
    Universe = function Universe() {
        if (instance) {
            return instance;
        }
        instance = this;
        // добавить остальную функциональность
        this.start_time = 0;
        this.bang = “Big”;
    };
}());
