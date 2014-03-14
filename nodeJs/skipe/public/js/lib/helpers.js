define(function () {
    return {
        validate: function (args) {
            if (_.isEmpty(args) && _.isArray(args)) {
                return;
            }
            var patterns = {
                email: /^.+@.+\..+$/,
                sname: /^[\w\s\_\-\=\+@]+$/,
            };
            var e = [];
            for (i in args) {
                if (!patterns[args[i].type].test(args[i].param)) {
                    e.push({param: args[i].key, msg: args[i].msg});
                }
            }
            return e;
        },
    };
});
