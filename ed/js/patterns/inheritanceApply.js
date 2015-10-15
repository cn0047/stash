var one = {
    name: "object",
    say: function (greet) {
        return greet + ", " + this.name;
    }
};
one.say('hi'); // "hi, object"
var two = {
    name: "another object"
};
one.say.apply(two, ['hello']); // "hello, another object"

var say = one.say;
say('hoho'); // "hoho, undefined"
// pass callback
var yetanother = {
    name: "Yet another object",
    method: function (callback) {
        return callback('Hola');
    }
};
yetanother.method(one.say); // "Hola, undefined"

function bind(o, m) {
    return function () {
        return m.apply(o, [].slice.call(arguments));
    };
}

var twosay = bind(two, one.say);
twosay('yo'); // "yo, another object"
