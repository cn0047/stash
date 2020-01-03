function* iter () {
 for (var i = 0; i < 10; i++) yield i;
}
for (var val of iter()) {
 console.log(val); // outputs 0 — 9
}
