var s = 'fswelxvda';
var s = ['f','s','w','e','l','x','v','d','a'];
var n = s.length;
for (i = 0; i < n; i++) {
    for (j = 0; j < n - i - 1; j++) {
        if (s[j] > s[j+1]) {
            var v = s[j];
            s[j] = s[j+1];
            s[j+1] = v;
        }
    }
}
