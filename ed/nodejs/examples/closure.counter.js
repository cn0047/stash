var c = (function () {
    var i = 0;
    return function () {
        return ++i;
    };
})();

c(); // 1
c(); // 2
c(); // 3
