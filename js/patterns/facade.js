var myevent = {
    // ...
    stop: function (e) {
        e.preventDefault();
        e.stopPropagation();
    }
    // ...
};
var myevent = {
    // ...
    stop: function (e) {
        // прочие броузеры
        if (typeof e.preventDefault === “function”) {
            e.preventDefault();
        }
        if (typeof e.stopPropagation === “function”) {
            e.stopPropagation();
        }
        // IE
        if (typeof e.returnValue === “boolean”) {
            e.returnValue = false;
        }
        if (typeof e.cancelBubble === “boolean”) {
            e.cancelBubble = true;
        }
    }
    // ...
};
