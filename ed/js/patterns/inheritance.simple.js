function Parent (name) {
    this.name = name || 'Adam';
}
Parent.prototype.say = function () {
    return this.name;
};
function Child (name) {
    Parent.apply(this, arguments);
}
Child.prototype = new Parent();

var kid = new Child('Patrick');
kid.name;           // 'Patrick'
kid.say();          // 'Patrick'
delete kid.name;
kid.say();          // 'Adam'
