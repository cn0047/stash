// Public property
var module = {
    myprop: 1,
    getProp: function () {
        return this.myprop;
    }
};

// Private property
function Gadget() {
    // private
    var name = ‘iPod’;
    // public
    this.getName = function () {
        return name;
    };
};

// static
Gadget.isShiny = function () {
    return "you bet";
};

// simple method
Gadget.prototype.setPrice = function (price) {
    this.price = price;
};

// Private methods to public scope
(function () {
    var astr = "[object Array]",
        toString = Object.prototype.toString;
    function isArray(a) {
        return toString.call(a) === astr;
    }
    function indexOf(haystack, needle) {
        var i = 0,
        max = haystack.length;
        for (; i < max; i += 1) {
            if (haystack[i] === needle) {
                return i;
            }
        }
        return -1;
    }
    myarray = {
        isArray: isArray,
        indexOf: indexOf,
        inArray: indexOf
    };
}());


// MODULE
MYAPP.utilities.array = (function () {
    return {
            inArray: function (needle, haystack) {
                // ...
            },
            isArray: function (a) {
                // ...
            }
    };
}());

// Module dependencies
MYAPP.utilities.module = (function (app, global) {
    //
}(MYAPP, this));
