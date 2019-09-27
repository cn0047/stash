const add = function (n1, n2) {
    const a = String(n1);
    const b = String(n2);
    const al = a.length;
    const bl = b.length;
    const n = al > bl ? al : bl;
    const ad = n - al;
    const bd = n - bl;
    let r = '';
    let addVal = 0;
    for (let i = 0; i < n; i++) {
        let av = parseInt(a[n - ad - i - 1]) || 0;
        let bv = parseInt(b[n - bd - i - 1]) || 0;
        let v = String(av + bv + addVal);
        if (v.length === 2) {
            addVal = parseInt(v[0]);
            v = v[1];
        } else {
            addVal = 0;
        }
        r = String(v) + String(r);
    }
    r = String(addVal) + String(r);
    return r.replace(/^[0]+/, '');
};

const mul = function (n1, n2) {
    const a = String(n1);
    const b = String(n2);
    const al = a.length;
    const bl = b.length;
    let sumUp = [];
    for (let bi = bl - 1; bi >= 0; bi--) {
        let addVal = '';
        // Add zero to end
        let r = (new Array(bl - bi)).join(0);
        for (let ai = al - 1; ai >= 0; ai--) {
            let av = parseInt(a[ai]);
            let bv = parseInt(b[bi]);
            let v = String((av * bv) + addVal);
            if (v.length === 2) {
                addVal = parseInt(v[0]);
                v = v[1]
            } else {
                addVal = 0;
            }
            r = String(v) + String(r);
        }
        r = String(addVal) + String(r);
        sumUp.push(r);
    }
    let res = sumUp.reduce(function (sum, v) {
      return add(sum, v);
    }, '0');
    return res;
};

console.log(add('0', '001') == 1);
console.log(add('937', '12188') == (937 + 12188));
// console.log(add('0109', '11122200898') == (0109 + 11122200898));
console.log(add(99, 99) == (99 + 99));
console.log(add(125, 1952) == (125 + 1952));

// console.log(add(2390525, 3021952) === (2390525 + 3021952));
// console.log(add(111222333, 999000222111000) === (111222333 + 999000222111000));
// console.log(mul(982, 101) === (982 * 101));
// console.log(mul(2, 99) === (2 * 99));
// console.log(mul(523, 798) === (523 * 798)); // 417354
