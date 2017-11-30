// task 3

function cmp(a, b) {
    return a - b;
}

function solution(A, B) {
    var n = A.length;
    var m = B.length;
    A.sort(cmp);
    B.sort(cmp);
    var i = 0;
    for (var j = 0; j < n; j++) {
        while (i < m - 1 && B[i] < A[j]) {
            i += 1;
        }
        if (A[j] === B[i]) return A[j];
    }
    return -1;
}

console.log(solution([1, 3, 2, 1],  [4, 2, 5, 3, 2]));
// console.log(solution([4, 2, 5, 3, 2], [1, 3, 2, 1]));
// console.log(solution([2, 1],  [3, 3]));
// console.log(solution([1, 3, 5, 7],  [0, 0, 1, 4, 9]));
// console.log(solution([3, 5, 7, 9],  [1, 2, 4, 7, 10]));
// console.log(solution([2, 1],  [1, 3, 3, 4, 5, 0]));
// console.log(solution([2, 1, 9, 0, 1, 2, 10000],  [3, 3, 10000]));
