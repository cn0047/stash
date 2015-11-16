// 1,1,2,3,5,8,13,21

function f1(n)
{
    a = [];
    a.push(1);
    a.push(1);
    for (i = 2; i < n; i++) {
        a.push(a[i-1] + a[i-2]);
    }
    return a[n-1];
}

function f2(n)
{
    p1 = 1;
    p2 = 1;
    if (n === 1 || n === 2) {
        return 1;
    }
    for (i = 2; i < n; i++) {
        v = p1 + p2;
        p1 = p2;
        p2 = v;
    }
    return v;
}
