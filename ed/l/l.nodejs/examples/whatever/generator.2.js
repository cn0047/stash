function* channel () {
  var name = yield 'hello, what is your name?'; // [1]
  return 'well hi there ' + name;
}
var gen = channel();
console.log(gen.next().value); // hello, what is your name? [2]
console.log(gen.next('billy')); // well hi there billy [3]
