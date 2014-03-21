define([], function () {
    return  Backbone.Model.extend({
        validate: function (a, o) {
            var e = [];
            var patterns = {
                email: /^.+@.+\..+$/,
                sname: /^[\w\s\_\-\=\+@]+$/,
                password: /^\w{10}$/,
            };
            for (k in a) {
                if ((k in this.checkOptions) && !patterns[this.checkOptions[k].type].test(a[k])) {
                    e.push({param: this.checkOptions[k].key, msg: this.checkOptions[k].msg});
                }
            }
            if (!_.isEmpty(e)) {
                return e;
            }
        },
    });
});