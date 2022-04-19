const r = (function() {
  return [
    (() => this.x).bind({ x: 'inner' })(),
    (() => this.x)()
  ]
}).call({ x: 'outer' });
console.log(r); // [ 'outer', 'outer' ]
