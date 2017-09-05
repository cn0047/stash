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
Child.prototype = Object.create(Parent.prototype);
console.log(kid.name);           // 'Patrick'
console.log(kid.say());          // 'Patrick'
delete kid.name;
console.log(kid.say());          // 'Adam'
