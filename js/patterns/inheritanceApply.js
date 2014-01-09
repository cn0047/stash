var one = {
    name: “object”,
    say: function (greet) {
        return greet + “, “ + this.name;
    }
};
// проверка
one.say(‘hi’); // “hi, object”

var two = {
    name: “another object”
};
one.say.apply(two, [‘hello’]); // “hello, another object”



// в случае присваивания функции переменной
// ссылка `this` будет указывать на глобальный объект
var say = one.say;
say(‘hoho’); // “hoho, undefined”
// передача в виде функции обратного вызова
var yetanother = {
    name: “Yet another object”,
    method: function (callback) {
        return callback(‘Hola’);
    }
};
yetanother.method(one.say); // “Hola, undefined”

function bind(o, m) {
    return function () {
        return m.apply(o, [].slice.call(arguments));
    };
}

var twosay = bind(two, one.say);
twosay(‘yo’); // “yo, another object”
