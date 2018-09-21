const f = (a) => {
  [a = 'default'] = [a];
  return a;
}

console.log(f(1)); // 1
console.log(f()); // default
