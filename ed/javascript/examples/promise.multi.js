function f(p, cb) {
  let ms = Math.floor((Math.random() * 1000) + 1);
  setTimeout(function() {
    cb('R:' + p + ' After:' + ms);
  }, ms);
}

// f(1, function (s) { console.log(s); });
// f(2, function (s) { console.log(s); });
// f(3, function (s) { console.log(s); });

// new Promise(function (res) {
//   f(4, function (d) {
//     res(d);
//   })
// }).then(function (d) {
//   console.log(d);
// });

let promises = [
  new Promise(function (resolve) { f(5, function (d) { resolve(d); }) }),
  new Promise(function (resolve) { f(6, function (d) { resolve(d); }) }),
  new Promise(function (resolve) { f(7, function (d) { resolve(d); }) })
];
Promise.all(promises).then(function (data) {
  console.log(data);
});
