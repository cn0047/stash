define([], function () {
    return  Backbone.Model.extend({
        hash: '',
        defaults: {
        },
        urlRoot: function () {
            return 'guest/'+this.hash;
        },
        validate: function (a, o) {
            var e = app.helpers.validate([
                {param: a.email, type: 'email', key: 'email', msg: 'Invalid email.'},
            ]);
            if (!_.isEmpty(e)) {
                return e;
            }
        },
    });
});
