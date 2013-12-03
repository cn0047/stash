var globalValue;
    exports.setGlobal = function(val) {
globalValue = val;
};
exports.returnGlobal = function() {
    console.log(global);
    return globalValue;
};
/*
// To require this script type next code:
var mod1 = require('./test.node.js');

// Then do:
mod1.setGlobal(34);
var val = mod1.returnGlobal();

*/