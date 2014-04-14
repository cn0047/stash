define([], function () {
    return  Backbone.skipeModel.extend({
        hash: '',
        defaults: {
        },
        checkOptions: {
            type: {type: 'match', pattern: /^(email|sname)$/, key: 'type', msg: app.nls.invalidLoginType},
            token: {type: 'email', key: 'token', msg: app.nls.invalidEmail},
            password: {type: 'password', key: 'password', msg: app.nls.invalidPassword},
        },
        urlRoot: function () {
            return 'guest/'+this.hash;
        },
    });
});
