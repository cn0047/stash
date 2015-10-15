function Universe() {
    var instance;
    Universe = function Universe() {
        return instance;
    };
    Universe.prototype = this;
    instance = new Universe();
    instance.constructor = Universe;
    instance.start_time = 0;
    instance.bang = "Big";
    return instance;
}

Universe.prototype.nothing = true; // true
var uni = new Universe();
Universe.prototype.everything = true; // true
var uni2 = new Universe();
uni === uni2; // true
uni.nothing && uni.everything && uni2.nothing && uni2.everything; // true
uni.bang; // "Big"
uni.constructor === Universe; // true

var Universe;
(function () {
    var instance;
    Universe = function Universe() {
        if (instance) {
            return instance;
        }
        instance = this;
        this.start_time = 0;
        this.bang = "Big";
    };
}());
