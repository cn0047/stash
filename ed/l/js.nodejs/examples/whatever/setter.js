const o = {};
Object.defineProperty(o, 'foo', {
    set: function(v) {
        console.log(v);
    }
});

o.foo = 'bar';
