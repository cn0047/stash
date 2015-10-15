function Parent (name) {
    this.name = name || 'Adam';
}
Parent.prototype.say = function () {
    return this.name;
};
function Child (name) {}
function inherit (C, P) {
    C.prototype = new P();
}
inherit(Child, Parent);

var kid = new Child();
kid.say(); // “Adam”

function Child2(a, c, b, d) {
    Parent.apply(this, arguments);
}