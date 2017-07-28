function foo() {
    return this.a;
}

console.log(foo()); // undefined

var obj = {a: 1};
console.log(foo.apply(obj)); // 1

var obj2 = {a: 'OBJ2'};
console.log(foo.apply(obj2)); // OBJ2

var hardBoundedFoo = function() {
    return foo.apply(obj, arguments);
};

console.log(hardBoundedFoo.apply(obj2)); // 1
