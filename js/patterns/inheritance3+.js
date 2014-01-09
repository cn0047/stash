// родительский конструктор
function Parent(name) {
    this.name = name || ‘Adam’;
}
// добавление дополнительной функциональности в прототип
Parent.prototype.say = function () {
    return this.name;
};
// дочерний конструктор
function Child(name) {
    Parent.apply(this, arguments);
}
Child.prototype = new Parent();

var kid = new Child(“Patrick”);
kid.name;           // “Patrick”
kid.say();          // “Patrick”
delete kid.name;
kid.say();          // “Adam”
