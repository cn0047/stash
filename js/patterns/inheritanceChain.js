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
var kid = new Child(“Patrick”);
kid.name;       // “Patrick”
typeof kid.say; // “undefined” - BECAUSE PROTOTYPE NOT USED.
