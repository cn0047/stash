var f = function() {
    return 'f';
}
var o = {
    f: function() {
        return 'of';
    }
};
var trick = function(f, o) {
    f = function() {
        return '=F';
    };
    o.f = function() {
        return '=OF';
    };
}
trick(f, o);
console.log(f());
console.log(o.f());

/*
Result:
f
=OF
*/
