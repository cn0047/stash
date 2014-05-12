var child = Object.create(parent, {
    // описание в соответствии со стандартом ECMA­ cript 5
    age: { value: 2 }
});
child.hasOwnProperty(“age”); // true
