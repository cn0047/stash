// родительский конструктор
function Parent(name) {
    this.name = name || ‘Adam’;
}
// добавление дополнительной функциональности в прототип
Parent.prototype.say = function () {
    return this.name;
};
// пустой дочерний конст
function Child(name) {}
// здесь происходит магия наследования
inherit(Child, Parent);
function inherit(C, P) {
    C.prototype = new P();
}

var kid = new Child();
kid.say(); // “Adam”




function Child2(a, c, b, d) {
    Parent.apply(this, arguments);
}