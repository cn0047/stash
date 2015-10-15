var child = Object.create(parent, {
    age: { value: 2 }
});
child.hasOwnProperty(“age”); // true
