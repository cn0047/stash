function solution(A) {
    var b;
    var max;
    while (A.length > 0) {
        // Split array into 2 parts: 1st part - variable `b`,
        // second - remainder from array `A`.
        var v = A.shift();
        // `b` don't contain sub-array it contain max values
        // for extracted elements from origin array `A`,
        // it helps to have better performance.
        b = b > v ? b : v;
        var maxFromA = Math.max.apply(Math, A);
        // maxFromA may contain value '-Infinity' which we have to exclude
        if (isFinite(maxFromA)) {
            var newMax = Math.abs(b - maxFromA);
            max = max > newMax ? max : newMax;
        }
    }
    return max;
}

console.log(solution([1, 3, -3]));
console.log(solution([4, 3, 2, 5, 1, 1]));
