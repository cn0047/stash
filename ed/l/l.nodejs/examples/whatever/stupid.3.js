const r = (function() {
  let f = this ? class g { } : class h { };
  return [
    typeof f,
    typeof h
  ];
})();
console.log(r); // [ 'function', 'undefined' ]
