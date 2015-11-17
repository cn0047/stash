var sum = function (a) {
    return function (b) {
        return a + b;
    }
}

console.log(sum(2)(3)); // 5
console.log(sum(5)(8)); // 13
