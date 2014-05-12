var obj = {
    value: 1,
    increment: function () {
        this.value += 1;
        return this;
    },
    add: function (v) {
        this.value += v;
        return this;
    },
    shout: function () {
        alert(this.value);
    }
};
// цепочка из вызовов методов
obj.increment().add(3).shout(); // 5



var Person = function (name) {
    this.name = name;
}.
method(‘getName’, function () {
    return this.name;
}).
method(‘setName’, function (name) {
    this.name = name;
    return this;
});
