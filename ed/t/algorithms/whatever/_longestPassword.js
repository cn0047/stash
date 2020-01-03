/**
 * Get longest valid password.
 *
 * It has to contain only alphanumerical characters (a−z, A−Z, 0−9);
 * There should be an even number of letters;
 * There should be an odd number of digits.
 *
 * @param  {string} str Input string with possible password separated by space.
 *
 * @return {number} Length of longest valid password.
 */
function f(str) {
  var arr = str.split(' ').filter(function (v, k, a) {
    var ok = /^[a-zA-z0-9]+$/.test(v);
    var d = v.match(/([0-9])/g) || [];
    var l = v.match(/([a-zA-z])/g) || [];
    var evenLetters = l.length % 2 === 0;
    var oddDigits = d.length % 2 === 1;
    return ok && evenLetters && oddDigits;
  });
  if (arr.length === 0) {
    return -1;
  }
  var r = arr.sort(function (a, b) {
    return b.length - a.length;
  })[0];
  return r.length;
}

console.log(7, f('test 5 a0A pass007 ?xy1'));
console.log(-1, f('a'));
console.log(1, f('0'));
console.log(3, f('asdf! 3ab qqqq adw3'));
