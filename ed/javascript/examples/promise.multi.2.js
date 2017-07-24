f = (next) => {
  const n = Math.random();
  if ((n * 10) > 7) {
    throw 'ERROR-1';
  }
  next(n);
}

const fail = (e) => {
    return 'Promise failed because ' + e;
}

const promises = [
  new Promise(function (resolve) {
    f(function (d) { resolve(d); })
  }).catch(fail),
  new Promise(function (resolve) {
    f(function (d) { resolve(d); })
  }).catch(fail),
  new Promise(function (resolve) {
    f(function (d) { resolve(d); })
  }).catch(fail),
];
Promise.all(promises).then(function (data) {
  console.log(data);
});
