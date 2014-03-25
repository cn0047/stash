define([], function () {
    return  Backbone.skipeModel.extend({
        hash: '',
        defaults: {
        },
        checkOptions: {
            type: {type: 'match', pattern: /^(email|sname)$/, key: 'type', msg: 'Invalid login type.'},
            token: {type: 'email', key: 'token', msg: 'Invalid email.'},
            password: {type: 'password', key: 'password', msg: 'Invalid password.'},
        },
        urlRoot: function () {
            return 'guest/'+this.hash;
        },
    });
});
