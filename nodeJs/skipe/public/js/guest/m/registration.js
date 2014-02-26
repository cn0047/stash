define([], function () {
    return  Backbone.Model.extend({
        hash: '',
        defaults: {
        },
        urlRoot: function () {
            return 'guest/'+this.hash;
        },
        validate: function (a, o) {
            var e = [];
            m = /^.+@.+\..+$/;
            if (!m.test(a.email)) {
                e.push({param: 'email', msg: 'Invalid email.'});
            }
            m = /^[\w\s\_\-\=\+@]+$/;
            if (!m.test(a.sname)) {
                e.push({param: 'sname', msg: 'Invalid screen name.'});
            }
            if (!_.isEmpty(e)) {
                return e;
            }
        },
    });
});
