// task 1

function solution(A) {
    if (A.length === 3) return -1;
    var max = 0;
    while (A.length > 0) {
      var v1 = A.shift();
      for (var i = 0; i < A.length - 1; i++) {
        for (var j = 1; j < A.length; j++) {
          if (i === j) continue;
          var v2 = A[i];
          var v3 = A[j];
          if (v1 + v2 > v3 && v2 + v3 > v1 && v1 + v3 > v2) {
            var newMax = v1 + v2 + v3;
            if (newMax > max) {
              max = newMax;
            }
          }
        }
      }
    }
    return max;
}

console.log(solution([10, 2, 5, 1, 8, 20]));
console.log(solution([5, 10, 18, 7, 8, 3]));
