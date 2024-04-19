function* iter () {
  for (var i = 0; i < 10; i++) yield i;
}
for (var val of iter()) {
  console.log(val);
}

/*
Result:
0
1
2
3
4
5
6
7
8
9
*/
