define([], function () {
    return  Backbone.Model.extend({
        hash: '',
        urlRoot: function () {
            return app.routers.app.scope+'/'+this.hash;
        },
        validate: function (a, o) {
            var e = [];
            var patterns = {
                match: '',
                email: /^.+@.+\..+$/,
                sname: /^[\w\s\_\-\=\+@]+$/,
                password: /^\w{10}$/,
            };
            for (k in a) {
                if (k in this.checkOptions) {
                    var pattern = '';
                    if (this.checkOptions[k].type in patterns) {
                        pattern = patterns[this.checkOptions[k].type];
                    }
                    if (this.checkOptions[k].type == 'match') {
                        pattern = this.checkOptions[k].pattern;
                    }
                    if (!pattern.test(a[k])) {
                        e.push({param: this.checkOptions[k].key, msg: this.checkOptions[k].msg});
                    }
                }
            }
            if (!_.isEmpty(e)) {
                return e;
            }
        },
    });
});