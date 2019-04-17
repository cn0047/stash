function f(s) {
  if (s.length === 0) {
    return -1;
  }
  if (s.length === 1) {
    return 0;
  }
  var l = s.length;
  var a = s.split('');
  var i = 1;
  while (a.length > 0) {
    if (a.shift() !== a.pop()) {
      return -1;
    }
    if (a.length === 1) {
      return i;
    }
    i++;
  }
  return -1;
}

function f2(s) {
  var l = s.length;
  if (l === 0 || l % 2 === 0) {
    return -1;
  }
  if (l === 1) {
    return 0;
  }
  var m = Math.floor(s.length / 2);
  var s1 = s.substring(0, m);
  var s2 = s.substring(m + 1, l).split('').reverse().join('');
  return s1 === s2 ? m : -1;
}

// console.log(f2(''));
// console.log(f2('x'));
console.log(f2('racecar'));
// console.log(f2('abba'));
