// The plan:
// start from begin of array,
// check two adjacent elements and move bigger to right,
// now last element in array is in sorted position,
// go to next iteration.
var a = ['f','s','w','e','l','x','v','d','a'];
var n = a.length;
for (i = 0; i < n; i++) {
    for (j = 0; j < n - i - 1 ; j++) {
        if (a[j] > a[j + 1]) {
            var v = a[j + 1];
            a[j + 1] = a[j];
            a[j] = v;
        }
    }
}
console.log(a);
