function Universe() {
    var instance = this;
    this.start_time = 0;
    this.bang = "Big";
    Universe = function () {
        return instance;
    };
}

var uni = new Universe();
var uni2 = new Universe();
uni === uni2; // true

Universe.prototype.nothing = true;
var uni = new Universe();
Universe.prototype.everything = true;
var uni2 = new Universe();

uni.nothing;        // true
uni2.nothing;       // true
uni.everything;     // undefined
uni2.everything;    // undefined
uni.constructor.name; // "Universe"
uni.constructor === Universe; // false
