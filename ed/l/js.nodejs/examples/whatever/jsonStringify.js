const o = {
  now: new Date(),
  undefined: undefined,
  func: function () { return "ok" }
}
console.log(
  o.func(), // ok
  JSON.stringify(o), // {"now":"2018-04-05T12:00:49.196Z"}
);
