const o = {
  n: Number(100020003000400050006),
}
const j = JSON.stringify(o);
console.log(o, j); // { n: 100020003000400050000 } '{"n":100020003000400050000}' ‼️
